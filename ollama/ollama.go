package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rokkit-exe/golly/models"
)

type Ollama struct {
	URL        string
	HTTPClient *http.Client
}

func NewOllama(host string, port string) *Ollama {
	return &Ollama{
		URL:        fmt.Sprintf("http://%s:%s/api/chat", host, port),
		HTTPClient: &http.Client{},
	}
}

func (o *Ollama) StreamChat(model string, messages []models.ChatMessage) (<-chan *models.ChatResponseChunk, <-chan error) {
	out := make(chan *models.ChatResponseChunk)
	errCh := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errCh)

		reqBody := models.ChatRequest{
			Model:    model,
			Stream:   true,
			Messages: messages,
		}

		jsonData, err := json.Marshal(reqBody)
		if err != nil {
			errCh <- err
			return
		}

		req, err := http.NewRequest("POST", o.URL, bytes.NewBuffer(jsonData))
		if err != nil {
			errCh <- err
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := o.HTTPClient.Do(req)
		if err != nil {
			errCh <- err
			return
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			var chunk models.ChatResponseChunk
			if err := json.Unmarshal([]byte(line), &chunk); err != nil {
				// You could choose to skip invalid lines or return error
				fmt.Printf("Skipping invalid line: %q\n", line)
				continue
			}

			out <- &chunk

			if chunk.Done {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			errCh <- err
			return
		}
	}()

	return out, errCh
}

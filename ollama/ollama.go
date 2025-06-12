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
		URL:        fmt.Sprintf("http://%s:%s", host, port),
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

		req, err := http.NewRequest("POST", o.URL+"/api/chat/", bytes.NewBuffer(jsonData))
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

func (o *Ollama) Create(name string, from string, system string) (models.CreateResponse, error) {
	reqBody := models.CreateRequest{
		Name:   name,
		From:   from,
		System: system,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return models.CreateResponse{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", o.URL+"/api/create", bytes.NewBuffer(jsonData))
	if err != nil {
		return models.CreateResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := o.HTTPClient.Do(req)
	if err != nil {
		return models.CreateResponse{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CreateResponse{}, fmt.Errorf("failed to create model, status code: %d", resp.StatusCode)
	}

	var response models.CreateResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return models.CreateResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	fmt.Printf("Model created successfully: %s\n", response.Status)

	return models.CreateResponse{}, nil
}

func (o *Ollama) List() (models.ListResponse, error) {
	req, err := http.NewRequest("GET", o.URL+"/api/tags", nil)
	if err != nil {
		return models.ListResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := o.HTTPClient.Do(req)
	if err != nil {
		return models.ListResponse{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.ListResponse{}, fmt.Errorf("failed to list models, status code: %d", resp.StatusCode)
	}

	var models models.ListResponse
	if err := json.NewDecoder(resp.Body).Decode(&models); err != nil {
		return models, fmt.Errorf("failed to decode response: %w", err)
	}

	return models, nil
}

func (o *Ollama) Delete(model string) error {
	reqBody := models.DeleteRequest{
		Model: model,
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}
	req, err := http.NewRequest("DELETE", o.URL+"/api/delete", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := o.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete model, status code: %d", resp.StatusCode)
	}

	return nil
}

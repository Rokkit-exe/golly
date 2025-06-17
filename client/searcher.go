package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Rokkit-exe/golly/models"
)

type Searcher struct {
	URL    string
	Client *http.Client
}

func NewSearcher(url string) *Searcher {
	return &Searcher{
		URL:    url,
		Client: &http.Client{},
	}
}

func (s *Searcher) Search(query string) (models.SearchResponse, error) {
	reqBody := models.SearchRequest{
		Q: query,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal("Error parsing reqBody")
		return models.SearchResponse{}, err
	}

	req, err := http.NewRequest("POST", s.URL+"/api/chat/", bytes.NewBuffer(jsonData))
	if err != nil {
		return models.SearchResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return models.SearchResponse{}, err
	}
	defer resp.Body.Close()

	var SearchResponse models.SearchResponse

	err = json.NewDecoder(resp.Body).Decode(&SearchResponse)
	if err != nil {
		return models.SearchResponse{}, err
	}

	return SearchResponse, nil
}

func (s *Searcher) GetWebResults(urls []string) ([]models.WebResult, error) {
	var webResults []models.WebResult
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			fmt.Println("Request error:", err)
			return []models.WebResult{}, err
		}
		defer resp.Body.Close()

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Read error:", err)
			return []models.WebResult{}, nil
		}
		webResults = append(webResults, models.WebResult{
			URL:     u,
			Content: string(body),
		})
	}
	return webResults, nil
}

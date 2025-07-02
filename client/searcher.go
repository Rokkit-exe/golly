package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
	params := url.Values{}
	params.Add("q", query)
	params.Add("format", "json")
	fullURL := fmt.Sprintf("%s/search?%s", s.URL, params.Encode())
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(nil))
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
			log.Fatalf("Request failed: %v", err)
		}
		defer resp.Body.Close()

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatalf("Failed to parse HTML: %v", err)
		}

		// Extract all visible text
		var sb strings.Builder
		doc.Find("body").Each(func(i int, s *goquery.Selection) {
			text := strings.TrimSpace(s.Text())
			if text != "" {
				sb.WriteString(text)
			}
		})
		webResults = append(webResults, models.WebResult{
			URL:     u,
			Content: sb.String(),
		})
	}
	return webResults, nil
}

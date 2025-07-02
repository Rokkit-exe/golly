package agent

import (
	"fmt"

	"github.com/Rokkit-exe/golly/client"
	"github.com/Rokkit-exe/golly/models"
	"github.com/Rokkit-exe/golly/ui"
)

type Agent struct {
	Ollama   client.Ollama
	Searcher client.Searcher
	Config   models.Config
	UI       ui.UI
}

func (a Agent) Search(query string) (<-chan *models.ChatResponseChunk, <-chan error) {
	a.UI.PrintStatus("Searching for: " + query)
	searchQuery := a.RefactorQuery(query)
	a.UI.PrintStatus("Refactored query: " + searchQuery)
	response, err := a.Searcher.Search(searchQuery)
	if err != nil {
		fmt.Println(err.Error())
	}

	a.UI.PrintStatus("Search completed. Processing results...")

	response.FilterResults()

	a.UI.PrintStatus("Found " + fmt.Sprint(len(response.GetUrls())) + " URLs.")
	urls := response.GetUrls()

	a.UI.PrintStatus("Fetching web results...")
	webResult, err := a.Searcher.GetWebResults(urls)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(webResult) == 0 {
		return nil, nil
	}

	a.UI.PrintStatus("Building answer from web results...")

	prompt := a.Config.SystemPrompts.AnswerBuilder
	queryBuilder := a.QueryBuilder(query, webResult)
	messages := []models.ChatMessage{
		{
			Role:    "system",
			Content: prompt,
		},
		{
			Role:    "user",
			Content: queryBuilder,
		},
	}
	return a.Ollama.StreamChat(a.Config.Model, messages)
}

func (a *Agent) RefactorQuery(query string) string {
	prompt := a.Config.SystemPrompts.RefactorQuery
	model := a.Config.Model
	searchQuery := a.Ollama.Generate(model, prompt, query)
	if searchQuery == "" {
		fmt.Println("No refactored query generated")
		searchQuery = query
	}
	fmt.Println("Refactored query:", searchQuery)

	return searchQuery
}

func (a *Agent) QueryBuilder(query string, results []models.WebResult) string {
	var fullQuery string
	fullQuery += "## Query: " + query + "\n\n"
	fullQuery += "## Results:\n\n"
	for _, result := range results {
		fullQuery += "url: " + result.URL + "\nContent: " + result.Content + "\n\n"
	}

	return fullQuery
}

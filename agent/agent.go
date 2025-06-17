package agent

import (
	"fmt"

	"github.com/Rokkit-exe/golly/client"
)

type Agent struct {
	Ollama   client.Ollama
	Searcher client.Searcher
}

func (a *Agent) Search(query string) (string, error) {
	response, err := a.Searcher.Search(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	response.FilterResults()
	urls := response.GetUrls()
	_, err = a.Searcher.GetWebResults(urls)

	return "", err

}

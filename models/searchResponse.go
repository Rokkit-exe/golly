package models

type SearchResponse struct {
	Query               string         `json:"query"`
	NumberOfResults     int            `json:"number_of_results"`
	Results             []SearchResult `json:"results"`
	Corrections         []string       `json:"corrections"`
	Infoboxes           []string       `json:"infoboxes"`
	Suggestions         []string       `json:"suggestions"`
	UnresponsiveEngines []string       `json:"unresponsive_engines"`
}

func (s *SearchResponse) FilterResults() {
	var results []SearchResult
	for _, r := range s.Results {
		if r.Score > 2 {
			results = append(results, r)
		}
	}
	s.Results = results
}

func (s SearchResponse) GetUrls() []string {
	var urls []string
	for _, r := range s.Results {
		urls = append(urls, r.URL)
	}
	return urls
}

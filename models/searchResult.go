package models

type SearchResult struct {
	URL           string   `json:"url"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	PublishedDate string   `json:"publishedDate"` // Use time.Time if you want to parse as time
	Thumbnail     string   `json:"thumbnail"`
	Engine        string   `json:"engine"`
	Template      string   `json:"template"`
	ParsedURL     []string `json:"parsed_url"`
	ImgSrc        string   `json:"img_src"`
	Priority      string   `json:"priority"`
	Engines       []string `json:"engines"`
	Positions     []int    `json:"positions"`
	Score         float64  `json:"score"`
	Category      string   `json:"category"`
}

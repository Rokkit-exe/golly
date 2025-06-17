package models

type SearchRequest struct {
	Q          string `json:"q"`
	Categories string `json:"categories,omitempty"`
	Engines    string `json:"engines,omitempty"`
	Language   string `json:"language,omitempty"`
	Page       int    `json:"page,omitempty"`
	TimeRange  string `json:"time_range,omitempty"`
	Format     string `json:"format,omitempty"`
}

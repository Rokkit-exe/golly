package models

type ChatRequest struct {
	Model    string        `json:"model"`    // The model to use for the ChatRequest
	Stream   bool          `json:"stream"`   // Whether to stream the response
	Messages []ChatMessage `json:"messages"` // The messages in the ChatRequest
}

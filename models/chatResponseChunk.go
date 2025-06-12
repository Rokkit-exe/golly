package models

type ChatResponseChunk struct {
	Model   string  `json:"model"`   // The model used for the ChatResponseChunk
	Message Message `json:"message"` // The message in the response ChatResponseChunk
	Done    bool    `json:"done"`    // Whether the response is complete
}

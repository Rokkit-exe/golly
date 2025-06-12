package models

type Message struct {
	Role    string `json:"role"`    // "user" or "assistant"
	Content string `json:"content"` // The content of the Message
}

package models

type ChatMessage struct {
	Role    string
	Content string
}

type GroqRequest struct {
	Messages []ChatMessage `json:"messages"`
	Model    string        `json:"model"`
}

type GroqResponse struct {
	Choices []struct {
		Message ChatMessage
	}
}

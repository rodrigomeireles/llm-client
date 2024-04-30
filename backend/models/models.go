package models

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
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

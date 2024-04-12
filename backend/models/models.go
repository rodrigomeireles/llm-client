package models

type ChatMessage struct {
	Role    string
	Content string
}

type ChatHistory struct {
	messages []ChatMessage
}

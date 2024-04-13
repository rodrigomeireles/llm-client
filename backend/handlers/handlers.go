package handlers

import (
	"github.com/rodrigomeireles/gpt-client/backend/models"
	"github.com/rodrigomeireles/gpt-client/web/templates"
	"net/http"
)

var history = []models.ChatMessage{}

func ChatClientHandler(w http.ResponseWriter, r *http.Request) {
	// Render the main page template with any dynamic data
	ctx := r.Context()
	err := templates.ChatClient("Leitãozinho").Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	// Render the history page template with any dynamic data
	ctx := r.Context()
	println(ctx)
	messages := []models.ChatMessage{
		{Role: "user", Content: "hello"},
		{Role: "assistant", Content: "bonjour"},
		{Role: "user", Content: "ohayou"},
	}
	err := templates.History(messages).Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

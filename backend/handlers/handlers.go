package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rodrigomeireles/gpt-client/backend/models"
	"github.com/rodrigomeireles/gpt-client/web/templates"
)

var history = []models.ChatMessage{{Role: "system", Content: "You are an AI assistant."}, {Role: "assistant", Content: "Hello."}}

func CallGroqModel(w http.ResponseWriter, system models.ChatMessage) {
	history = append([]models.ChatMessage{system}, history...)
	bbody, _ := json.Marshal(history)
	res, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewReader(bbody))
	if err := templates.History(newMessage).Render(ctx, w); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func ChatClientHandler(w http.ResponseWriter, r *http.Request) {
	// Render the main page template with any dynamic data
	ctx := r.Context()
	err := templates.ChatClient("Leitãozinho").Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	// Render the history page template with any dynamic data
	ctx := r.Context()
	switch r.Method {
	case "GET":
		err := templates.History(history).Render(ctx, w)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

	case "POST":
		userMessage := r.PostFormValue("user_message")
		if userMessage == "" {
			http.Error(w, "Bad Request: no message provided", 400)
			return
		}

		newMessage := []models.ChatMessage{{Role: "user", Content: userMessage}}
		history = append(history, newMessage...)

		if err := templates.History(newMessage).Render(ctx, w); err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}
}

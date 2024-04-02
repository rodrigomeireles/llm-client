package handlers

import (
	"github.com/rodrigomeireles/gpt-client/web/templates"
	"net/http"
)

func ChatClientHandler(w http.ResponseWriter, r *http.Request) {
	// Render the main page template with any dynamic data
	ctx := r.Context()
	err := templates.ChatClient("Leitãozinho").Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {

	// Render the main page template with any dynamic data
	ctx := r.Context()
	err := templates.HistoryWindow().Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

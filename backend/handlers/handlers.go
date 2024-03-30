package handlers

import (
	"fmt"
	"github.com/rodrigomeireles/gpt-client/web/templates"
	"html"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	// Render the main page template with any dynamic data
	ctx := r.Context()
	err := templates.ChatClient("Leitãozinho").Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	if err != nil {
		http.Error(w, "History handler error", 500)
	}
}

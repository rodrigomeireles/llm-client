package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/rodrigomeireles/gpt-client/backend/models"
	"github.com/rodrigomeireles/gpt-client/web/templates"
)

var history = []models.ChatMessage{{Role: "system", Content: "You are an AI assistant."}}

func CallGroqModel(messages *[]models.ChatMessage) (*http.Response, error) {
	// form the request
	// this assumes the API-KEY was already loaded as an environment variable
	body := &models.GroqRequest{
		Model:    "mixtral-8x7b-32768",
		Messages: *messages,
	}
	bbody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewReader(bbody))
	if err != nil {
		return nil, err
	}
	key := os.Getenv("GROQ_API_KEY")
	if key == "" {
		return nil, errors.New("Empty API-KEY")
	}
	req.Header.Set("Authorization", "Bearer "+key)
	//call the model
	client := &http.Client{}
	return client.Do(req)
}

func ChatClientHandler(w http.ResponseWriter, r *http.Request) {
	// Render the main page template with any dynamic data
	ctx := r.Context()
	err := templates.ChatClient("Leitãozinho").Render(ctx, w)
	if err != nil {
		log.Panic(err)
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
			log.Println("Empty message")
			return
		}
		newMessage := []models.ChatMessage{{Role: "user", Content: userMessage}}
		history = append(history, newMessage...)
		llm_res, err := CallGroqModel(&history)
		if err != nil {
			http.Error(w, "Error calling the model", 502)
			log.Println(err)
			return
		}
		// map Groq response and Unmarshall it
		var res models.GroqResponse
		resbytes, _ := io.ReadAll(llm_res.Body)
		_ = json.Unmarshal(resbytes, &res)
		history = append(history, []models.ChatMessage{res.Choices[0].Message}...)
		log.Println(history)
		err = templates.History(history).Render(ctx, w)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}
}

func SidebarHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := templates.Sidebar().Render(ctx, w)
	if err != nil {
		log.Panic(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

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

var history = []models.ChatMessage{{Role: "system", Content: "You are an AI assistant with a snarky attitude."}}

func CallGroqModel(messages *[]models.ChatMessage, model string) (*http.Response, error) {
	// form the request
	// this assumes the API-KEY was already loaded as an environment variable
	body := &models.GroqRequest{
		Model:    model,
		Messages: *messages,
	}
	bbody, _ := json.Marshal(body)
	log.Printf("Calling model with body %s", string(bbody))
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

func GetHistoryHandler(w http.ResponseWriter, r *http.Request) {
	// Render the history page template with any dynamic data
	ctx := r.Context()
	err := templates.History(history).Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func PostHistoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userMessage := r.PostFormValue("user_message")
	model := r.PostFormValue("model")
	if userMessage == "" {
		http.Error(w, "Bad Request: no message provided", 400)
		log.Println("Empty message")
		return
	}
	newMessage := []models.ChatMessage{{Role: "user", Content: userMessage}}
	history = append(history, newMessage...)
	log.Println("Calling model!")
	llm_res, err := CallGroqModel(&history, model)
	if err != nil {
		http.Error(w, "Error calling the model", 502)
		log.Printf("Error calling the model: %v", err)
		return
	}
	log.Println("Reading response to byte-array")
	resbytes, err := io.ReadAll(llm_res.Body)

	if err != nil {
		log.Println("Jesus Christ we can't read shit")
		http.Error(w, "Error decoding the response", http.StatusUnprocessableEntity)
		return
	}

	if llm_res.StatusCode != 200 {
		http.Error(w, "Error calling the model", 502)
		log.Printf("Error calling the model:\n Response Code: %d, Response: %s", llm_res.StatusCode, string(resbytes))
		return
	}
	// map Groq response and Unmarshall it
	var res models.GroqResponse
	log.Println("Unmarshalling response")
	marsh_err := json.Unmarshal(resbytes, &res)
	if marsh_err != nil {
		http.Error(w, "Error decoding the response", http.StatusUnprocessableEntity)
		log.Printf("Error unmarshalling the response: %v", res)
		return
	}
	history = append(history, []models.ChatMessage{res.Choices[0].Message}...)
	log.Println(history)
	err = templates.History(history).Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

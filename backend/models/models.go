package models

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqModel int

const (
	LLaMA3_8b_8192 GroqModel = iota
	LLaMA3_70b_8192
	Mixtral_8x7b
	Gemma_7b
	Gemma_9b
)

var GroqModels = map[string]string{
	LLaMA3_8b_8192.String():  "llama3-8b-8192",
	LLaMA3_70b_8192.String(): "llama3-70b-8192",
	Mixtral_8x7b.String():    "mixtral-8x7b-32768",
	Gemma_7b.String():        "gemma2-7b-it",
	Gemma_9b.String():        "gemma2-9b-it",
}

type GroqRequest struct {
	Messages    []ChatMessage `json:"messages"`
	Model       string        `json:"model"`
	Temperature float64       `json:"temperature"`
	Top_p       float64       `json:"top_p"`
}

type GroqResponse struct {
	Choices []struct {
		Message ChatMessage
	}
}

type Config struct {
	Model       string  `json:"model"`
	Temperature float64 `json:"temperature"`
	Top_p       float64 `json:"top_p"`
}

type ClientState struct {
	History []ChatMessage `json:"history"`
	Config  Config        `json:"-"`
}

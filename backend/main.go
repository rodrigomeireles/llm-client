package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/rodrigomeireles/llm-client/backend/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	rootdir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error while aquiring rootdir.")
	}
	log.Println(rootdir)
	http.Handle("GET /images/", http.StripPrefix("/images",
		http.FileServer(http.Dir(path.Join(rootdir, "web/images")))))

	http.Handle("GET /styles/", http.StripPrefix("/styles",
		http.FileServer(http.Dir(path.Join(rootdir, "web/static/css")))))

	http.Handle("GET /scripts/", http.StripPrefix("/scripts",
		http.FileServer(http.Dir(path.Join(rootdir, "scripts")))))

	http.Handle("GET /history",
		http.HandlerFunc(handlers.GetHistoryHandler))

	http.Handle("POST /history",
		http.HandlerFunc(handlers.PostHistoryHandler))

	http.HandleFunc("/", handlers.ChatClientHandler)

	log.Printf("Starting server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

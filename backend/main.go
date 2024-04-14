package main

import (
	"github.com/joho/godotenv"
	"github.com/rodrigomeireles/gpt-client/backend/handlers"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rootdir, err := os.Getwd()
	if err != nil {
		rootdir = "woops"
	}
	http.Handle("/images/", http.StripPrefix("/images",
		http.FileServer(http.Dir(path.Join(rootdir, "web/images")))))

	http.Handle("/styles/", http.StripPrefix("/styles",
		http.FileServer(http.Dir(path.Join(rootdir, "web/static/css")))))

	http.Handle("/scripts/", http.StripPrefix("/scripts",
		http.FileServer(http.Dir(path.Join(rootdir, "scripts")))))

	http.Handle("/history",
		http.HandlerFunc(handlers.HistoryHandler))

	http.HandleFunc("/", handlers.ChatClientHandler)

	log.Printf("Starting server.")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

#!/bin/bash

# Start Tailwind CSS in watch mode
echo "Starting Tailwind CSS in watch mode..."
npx tailwindcss -i ./web/static/css/styles.css -o ./web/static/css/output.css --watch=always &
# Save the Tailwind process PID
TAILWIND_PID=$!


# Start the Go application in the background
echo "Starting Go server and Templ watcher with live reload..."
templ generate --watch --proxy="http://localhost:8080" --cmd="go run ./backend/main.go" &

# Save the Go application PID
GO_PID=$!


# Function to kill the processes on exit
cleanup() {
    echo "Stopping Go server and Tailwind..."
    kill $GO_PID
    pkill main
    kill $TAILWIND_PID
    exit
}

# Trap script exit and execute the cleanup function
trap cleanup SIGINT

# Wait for the processes to end
wait $GO_PID
wait $TAILWIND_PID

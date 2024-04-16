#!/bin/bash

# Navigate to the project backend directory
cd backend

# Build the Go application
echo "Building Go server..."
go build -o cvserver.bin
echo "Build complete"

cd ..
# Start the Go application
./backend/cvserver.bin

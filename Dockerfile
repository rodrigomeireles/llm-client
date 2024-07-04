FROM node:latest AS node_builder
COPY /web /web
RUN npm install -D tailwindcss && npx tailwindcss -i ./web/static/css/styles.css -o ./web/static/css/output.css

# Fetch
FROM golang:latest AS fetch-stage
COPY go.mod go.sum /app
WORKDIR /app
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Build
FROM golang:latest AS build-stage
COPY --from=generate-stage /app /app
WORKDIR /app/backend
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false

# Test
FROM build-stage AS test-stage
RUN go test -v ./...

# Deploy
FROM alpine:latest AS deploy-stage
COPY --chown=65532:65532 --from=build-stage /app /app
WORKDIR /app
EXPOSE 8080
CMD [".backend/backend"]

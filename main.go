// DeepSeek Web Chat - Simple Golang Server
// ----------------------------------------
// A lightweight web server that connects to DeepSeek API using OpenAI-compatible interface.
// This server receives a prompt from the frontend, sends it to DeepSeek, and returns the response.
// Includes a system prompt as a marketing assistant for Westown View apartment.
//
// Structure:
// - Loads API key from .env
// - Serves static frontend (index.html)
// - Handles POST /chat endpoint for chatting
//
// Author: Kamu (dibantu Freya)
// Requirements: go 1.18+, ngrok, API Key from https://platform.deepseek.com

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

// PromptRequest is the expected JSON body from the frontend.
type PromptRequest struct {
	Prompt string `json:"prompt"`
}

// ResponseData is the JSON format returned to the frontend.
type ResponseData struct {
	Answer string `json:"answer"`
}

// handleChat processes POST /chat requests and connects to DeepSeek API.
func handleChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST only", http.StatusMethodNotAllowed)
		return
	}

	var input PromptRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Load API Key from environment
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		http.Error(w, "Missing API Key", http.StatusInternalServerError)
		return
	}

	// Create DeepSeek client using OpenAI-compatible SDK
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.deepseek.com/v1"
	client := openai.NewClientWithConfig(config)

	// Set system prompt: Asisten Westown View
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "deepseek-chat",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "Kamu adalah asisten marketing dari Westown View. Tugasmu adalah membantu menjelaskan apartemen Westown View kepada calon customer dengan bahasa yang sopan, informatif, dan meyakinkan.",
			},
			{
				Role:    "user",
				Content: input.Prompt,
			},
		},
	})

	if err != nil {
		http.Error(w, "Chat error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response back as JSON
	json.NewEncoder(w).Encode(ResponseData{Answer: resp.Choices[0].Message.Content})
}

// main initializes the server, loads .env, and sets up routing.
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system env")
	}

	// Serve static frontend files from ./static
	http.Handle("/", http.FileServer(http.Dir("static")))

	// Handle chat endpoint
	http.HandleFunc("/chat", handleChat)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

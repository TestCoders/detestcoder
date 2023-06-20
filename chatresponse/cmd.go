package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Struct to represent the API request payload
type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIRequest struct {
	Model     string          `json:"model"`
	Messages  []OpenAIMessage `json:"messages"`
	MaxTokens string          `json:"max_tokens"`
}

// Struct to represent the API response payload
type GPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	// Parse command-line arguments
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Please provide a prompt.")
		os.Exit(1)
	}

	// Construct the API request payload
	requestPayload := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []OpenAIMessage{
			{Role: "user", Content: args[0]},
		},
		MaxTokens: "32",
	}

	// Convert the request payload to JSON
	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		os.Exit(1)
	}

	// Create an HTTP client and send the request to the ChatGPT API
	client := http.Client{
		Timeout: time.Second * 10,
	}
	apiURL := "https://api.openai.com/v1/chat/completions"
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating API request:", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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") // Replace with your actual API key

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending API request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the API response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading API response:", err)
		os.Exit(1)
	}

	// Print the entire API response
	fmt.Println(string(responseBody))
}

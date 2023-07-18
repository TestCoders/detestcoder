/*
Copyright © 2023 TestCoders / DeTesters
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/resty.v1"
)

type OpenAIRequest struct {
	Model     string          `json:"model"`
	Messages  []OpenAIMessage `json:"messages"`
	MaxTokens string          `json:"max_tokens"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Model  string `json:"model"`
}

func main() {
	// Set your OpenAI API key
	apiKey := "INPUT API KEY"

	// Set the API endpoint
	apiURL := "https://api.openai.com/v1/chat/completions"

	// Prepare the request payload
	requestData := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []OpenAIMessage{
			{Role: "user", Content: "Hello!"},
		},
		MaxTokens: "32",
	}

	// Create a new Resty client
	client := resty.New()

	// Set the authorization header with your API key
	client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Send the API request
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestData).
		Post(apiURL)

	if err != nil {
		log.Fatalf("Error making API request: %v", err)
	}

	// Parse the API response
	var response OpenAIResponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		log.Fatalf("Error parsing API response: %v", err)
	}

	// Display the completion text
	fmt.Println("Generated Text:", response.ID)
	fmt.Println("Status:", resp.Status())
	fmt.Println("Body", string(resp.Body()))

}

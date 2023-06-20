/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "chatgpt",
	Short: "Commandline prompter for chatgpt",
	Long: `This is a commandline tool to send prompts and receive responses from chatGPT
	
	Make sure you have set an environment variable OPENAI_API_KEY containing your API key 
	and provide a prompt as a commandline argument`,
	Run: runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set the OPENAI_API_KEY environment variable.")
		os.Exit(1)
	}

	if len(args) == 0 {
		fmt.Println("Please provide a prompt.")
		os.Exit(1)
	}
	prompt := args[0]

	// Construct the API request payload
	requestPayload := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []OpenAIMessage{
			{Role: "user", Content: prompt},
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
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending API request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the API response
	// responseBody := ioutil.ReadAll
	// fmt.Println("response", responseBody)
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

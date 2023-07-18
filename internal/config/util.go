package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var models = []string{"OpenAI"}

type UserConfig struct {
	ApiKey       string
	Model        string
	ModelVersion string
}

type configPrompt struct {
	label    string
	errorMsg string
}

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

// ReadConfig loads the .detestcoder.yaml config file from the users' home directory
func ReadConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".detestcoder")
	viper.AutomaticEnv()

	cobra.CheckErr(viper.ReadInConfig())
	fmt.Println(viper.Get("api_key"))
}

// WriteConfig writes the chosen settings to a .detestcoder.yaml file in the users' home directory
// NOTE: any subsequent call to this will overwrite existing settings
func WriteConfig() error {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	config := getUserInput()

	// Configure Viper
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".detestcoder")

	// Set config values
	viper.Set("model", config.Model)
	viper.Set("model_version", config.ModelVersion)
	viper.Set("api_key", config.ApiKey)

	// Workaround for creating a config file if it doesn't exist
	if err := viper.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			return viper.WriteConfig()
		}
	}

	return nil
}

// getUserInput uses an interactive prompt to retrieve settings input from the user
func getUserInput() UserConfig {
	modelPrompt := configPrompt{
		label:    "Which model do you want to use? ",
		errorMsg: "Select a model",
	}
	model := getUserInputSelect(modelPrompt, models)

	modelVersionPrompt := configPrompt{
		label:    fmt.Sprintf("Which version of '%v' would you like to use? ", model),
		errorMsg: "Provide a version",
	}
	modelVersion := getUserInputString(modelVersionPrompt, true, false)

	apiKeyPrompt := configPrompt{
		label: fmt.Sprintf("Provide your '%v' API key ", model),
	}
	apiKey := getUserInputString(apiKeyPrompt, false, true)

	return UserConfig{
		ApiKey:       apiKey,
		Model:        model,
		ModelVersion: modelVersion,
	}
}

// getUserInputSelect creates a prompt where the user can provide textual input.
func getUserInputString(cp configPrompt, allowEmpty, mask bool) string {
	validate := func(input string) error {
		if !allowEmpty && len(input) <= 0 {
			return errors.New(cp.errorMsg)
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:          "{{ . }}",
		Confirm:         "{{ . }}",
		Valid:           "{{ . | green }}",
		Invalid:         "{{ . | red }}",
		Success:         "{{ . | green}}",
		ValidationError: "{{ . | red }}",
		FuncMap:         nil,
	}

	var maskRune rune

	if mask {
		maskRune = '*'
	} else {
		maskRune = 0
	}

	prompt := promptui.Prompt{
		Label:     cp.label,
		Templates: templates,
		Validate:  validate,
		Mask:      maskRune,
	}

	result, err := prompt.Run()
	cobra.CheckErr(err) // NOTE: use own check err?
	return result
}

// getUserInputSelect creates a prompt where the user can select any of the provided items
func getUserInputSelect(cp configPrompt, items []string) string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label: cp.label,
			Items: items,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	cobra.CheckErr(err)
	return result
}

func PromptInit() error {
	ReadConfig()
	apiKey := viper.GetString("api_key")

	promptInput := configPrompt{
		label:    fmt.Sprintf("Input your prompt "),
		errorMsg: "Provide a version",
	}

	prompt := getUserInputString(promptInput, false, false)

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
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading API response:", err)
		os.Exit(1)
	}

	// Print the entire API response
	fmt.Println(string(responseBody))
	return nil
}

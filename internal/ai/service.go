package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/testcoders/detestcoder/pkg/config/aimodel"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIService struct {
	Client      HttpClient
	APIKey      string
	Model       string
	Messages    []Message
	Temperature float64
	Stream      bool
}

type MockOpenAIService struct {
	Client      HttpClient
	APIKey      string
	Model       string
	Messages    []Message
	Temperature float64
	Stream      bool
}

// Service interface
type Service interface {
	Send() ([]byte, error)
}

type ServiceStore struct{}

func NewService() *ServiceStore {
	return &ServiceStore{}
}

func (ss *ServiceStore) GetService(prompt string, model aimodel.AIModel) Service {
	switch model.AiModel {
	case "OpenAI":
		return &OpenAIService{
			Client: &RealHttpClient{},
			APIKey: model.ApiKey,
			Model:  model.ModelVersion,
			Messages: []Message{
				{
					Role:    "system",
					Content: prompt,
				},
			},
			Temperature: 0,
			Stream:      false,
		}
	case "Mock":
		return &MockOpenAIService{
			Client: &RealHttpClient{},
			APIKey: model.ApiKey,
			Model:  model.ModelVersion,
			Messages: []Message{
				{
					Role:    "system",
					Content: prompt,
				},
			},
			Temperature: 0,
			Stream:      false,
		}
	default:
		return nil
	}
}

type Prompt struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	Stream      bool      `json:"stream"`
}

func (s *OpenAIService) Send() ([]byte, error) {
	openaiURL := "https://api.openai.com/v1/chat/completions"

	promptStruct := &Prompt{
		Model:       s.Model,
		Messages:    s.Messages,
		Temperature: s.Temperature,
		Stream:      s.Stream,
	}

	jsonPrompt, err := json.Marshal(promptStruct)
	if err != nil {
		log.Printf("Failed to marshal prompt structure: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(jsonPrompt))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.APIKey))

	//Print request details
	fmt.Printf("\nRequest details:\n")
	fmt.Printf("Method: %v\n", req.Method)
	fmt.Printf("URL: %v\n", req.URL)
	fmt.Printf("Headers: %v\n", req.Header)
	fmt.Printf("Body: %s\n", jsonPrompt)

	resp, err := s.Client.Do(req)
	if err != nil {
		log.Printf("Failed to execute request: %v", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	//Print response details
	fmt.Printf("\nResponse details:\n")
	fmt.Printf("Status: %v\n", resp.Status)
	fmt.Printf("Headers: %v\n", resp.Header)
	fmt.Printf("Body: %s", body)

	return body, nil
}

func (s *MockOpenAIService) Send() ([]byte, error) {
	openAIURL := "https://api.openai.com/v1/chat/completions"

	promptStruct := &Prompt{
		Model:       s.Model,
		Messages:    s.Messages,
		Temperature: s.Temperature,
		Stream:      s.Stream,
	}

	jsonPrompt, err := json.Marshal(promptStruct)
	if err != nil {
		log.Printf("Failed to marshal prompt structure: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", openAIURL, bytes.NewBuffer(jsonPrompt))
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.APIKey))

	// Here `s.Client.Do(req)` is replaced by `s.Do(req)`
	resp, err := s.Do(req)
	if err != nil {
		log.Printf("Failed to execute request: %v", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	return body, nil
}

package ai

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

// HttpClient interface
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// RealHttpClient Real http client that will make the requests
type RealHttpClient struct{}

func (h *RealHttpClient) Do(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

func (s *MockOpenAIService) Do(req *http.Request) (*http.Response, error) {
	mockJson, err := os.ReadFile("mockdata/mock_response.json")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	resp := io.NopCloser(bytes.NewReader(mockJson))
	return &http.Response{
		StatusCode: 200,
		Body:       resp,
	}, nil
}

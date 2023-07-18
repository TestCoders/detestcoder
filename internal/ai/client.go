package ai

import (
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Service interface describes the methods an AI backend should implement
type Service interface {
	Send(prompt string) (*Response, error)
}

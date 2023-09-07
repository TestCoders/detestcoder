package ai

import (
	"encoding/json"
	"github.com/testcoders/detestcoder/pkg/config/aimodel"
)

// The Response interface
type Response struct {
	Created int64  `json:"created"`
	Content string `json:"content"`
	Role    string `json:"role"`
	// We can add more, if we think we need it (like Usage tokens, for calculating costs)
}

// OpenAIResponse is a generic way to handle different AI backends' response data
type OpenAIResponse struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Message      struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
	} `json:"choices"`
	Created int64  `json:"created"`
	ID      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func (r *Response) GetResponse(raw []byte, model aimodel.AIModel) Response {
	switch model.AiModel {
	case "OpenAI":
		openaiResponse := new(OpenAIResponse)
		err := json.Unmarshal(raw, openaiResponse)
		if err != nil {
			return Response{}
		}
		r.Created = openaiResponse.Created
		r.Content = openaiResponse.Choices[0].Message.Content
		r.Role = openaiResponse.Choices[0].Message.Role
	case "Mock":
		openaiResponse := new(OpenAIResponse)
		err := json.Unmarshal(raw, openaiResponse)
		if err != nil {
			return Response{}
		}
		r.Created = openaiResponse.Created
		r.Content = openaiResponse.Choices[0].Message.Content
		r.Role = openaiResponse.Choices[0].Message.Role
	}
	return Response{}
}

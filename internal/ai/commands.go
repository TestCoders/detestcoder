package ai

import "encoding/json"

// SendPrompt is used to send the prompt to any AI backend
func SendPrompt(service Service) (*Response, error) {
	// Do additional things that are backend independent here, for example
	// validate initialize or output text to the terminal

	raw, err := service.Send()
	if err != nil {
		return nil, err
	}

	response := new(Response)
	err = json.Unmarshal(raw, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

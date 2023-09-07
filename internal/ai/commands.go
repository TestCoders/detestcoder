package ai

import (
	"github.com/testcoders/detestcoder/pkg/config/aimodel"
)

// SendPrompt is used to send the prompt to any AI backend and giving back a Response object
func SendPrompt(service Service, model aimodel.AIModel, verbose bool) (*Response, error) {
	raw, err := service.Send(verbose)
	if err != nil {
		return nil, err
	}

	response := new(Response)
	response.GetResponse(raw, model)

	return response, nil
}

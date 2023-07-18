package ai

// SendPrompt is used to send the prompt to any AI backend
func SendPrompt(service Service, prompt string) (*Response, error) {

	// Do additional things that are backend independent here, for example
	// validate config or output text to the terminal

	return service.Send(prompt)
}

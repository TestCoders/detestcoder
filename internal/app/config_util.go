package app

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
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

// ReadConfig loads the .detestcoder.yaml config file from the users' home directory
func ReadConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".detestcoder")
	viper.AutomaticEnv()

	cobra.CheckErr(viper.ReadInConfig())
}

// WriteConfig writes the chosen settings to a .detestcoder.yaml file in the users' home directory
// NOTE: any subsequent call to this will overwrite existing settings
func WriteConfig() {
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
			cobra.CheckErr(viper.WriteConfig())
		}
	}
}

// getUserInput uses an interactive prompt to retrieve settings input from the user
func getUserInput() UserConfig {
	modelPrompt := configPrompt{
		label:    "Which model do you want to use?",
		errorMsg: "Select a model",
	}
	model := getUserInputSelect(modelPrompt, models)

	modelVersionPrompt := configPrompt{
		label:    fmt.Sprintf("Which version of '%v' would you like to use?", model),
		errorMsg: "Provide a version",
	}
	modelVersion := getUserInputString(modelVersionPrompt, true)

	apiKeyPrompt := configPrompt{
		label: fmt.Sprintf("Provide your '%v' API key ", model),
	}
	apiKey := getUserInputString(apiKeyPrompt, false)

	return UserConfig{
		ApiKey:       apiKey,
		Model:        model,
		ModelVersion: modelVersion,
	}
}

// getUserInputSelect creates a prompt where the user can provide textual input.
func getUserInputString(cp configPrompt, allowEmpty bool) string {
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

	prompt := promptui.Prompt{
		Label:     cp.label,
		Templates: templates,
		Validate:  validate,
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

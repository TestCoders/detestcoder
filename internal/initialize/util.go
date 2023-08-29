package initialize

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/testcoders/detestcoder/pkg/config"
	"github.com/testcoders/detestcoder/pkg/config/aimodel"
	"os"
	"path"
)

var models = []string{"OpenAI"}
var modelVersions = []string{"gpt-4", "gpt-3.5-turbo"}
var yn = []string{"y", "n"}

// ReadConfig loads the .detestcoder initialize file from the projects' working directory
func ReadConfig() (*aimodel.AIModel, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(homeDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".detestcoder")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var aiModel aimodel.AIModel
	err = viper.Unmarshal(&aiModel)
	if err != nil {
		return nil, err
	}

	return &aiModel, nil
}

// WriteConfig writes the chosen settings to a .detestcoder.yaml file in the projects' working directory
// NOTE: any subsequent call to this will overwrite existing settings
func WriteConfig() error {
	if isDetestCoderInitialized() {
		if askUserToProceed() == "n" {
			os.Exit(0)
		}
	}

	aiModel := aimodel.NewAiModel()

	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	getAiModelInput(aiModel)

	// Configure Viper
	viper.AddConfigPath(homeDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".detestcoder")

	// Set initialize values
	viper.Set("AIModel", aiModel)

	cfgFile := path.Join(homeDir, ".detestcoder.yaml")

	// Create the file if it does not exist.
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		_, err = os.Create(cfgFile)
		if err != nil {
			return err
		}
	}

	// Overwrite the existing initialize or Write a new one
	err = viper.WriteConfigAs(cfgFile)
	if err != nil {
		return err
	}

	return nil
}

// askUserToProceed checks whether the user wants to overwrite the existing setup
func askUserToProceed() string {
	proceedPrompt := config.ConfigPrompt{
		Label:    "It seems detestcoder is already initialized. You can update the existing setup this way. Do you wish to continue? ",
		ErrorMsg: "Select 'y' or 'n'.",
	}
	return config.GetUserInputSelect(proceedPrompt, yn)
}

// isDetestCoderInitialized checks whether detestcoders is already initialized
func isDetestCoderInitialized() bool {
	workingDir, err := os.Getwd()
	if err != nil {
		return false
	}

	configPath := workingDir + "/.detestcoder.yaml"
	if _, err := os.Stat(configPath); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

// getAiModelInput uses an interactive prompt to retrieve settings input from the user
func getAiModelInput(aiModel *aimodel.AIModel) {
	modelNamePrompt := config.ConfigPrompt{
		Label:    "Which aimodel do you want to use? ",
		ErrorMsg: "Select an aimodel",
	}
	modelName := config.GetUserInputSelect(modelNamePrompt, models)
	aiModel.SetModel(modelName)

	modelVersionPrompt := config.ConfigPrompt{
		Label:    fmt.Sprintf("Which version of '%v' would you like to use? ", modelName),
		ErrorMsg: "Provide a version",
	}
	modelVersion := config.GetUserInputSelect(modelVersionPrompt, modelVersions)
	aiModel.SetModelVersion(modelVersion)

	apiKeyPrompt := config.ConfigPrompt{
		Label: fmt.Sprintf("Provide your '%v' API key ", modelName),
	}
	apiKey := config.GetUserInputString(apiKeyPrompt, false, true)
	aiModel.SetApiKey(apiKey)
}

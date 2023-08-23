package config

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/testcoders/detestcoder/pkg/config/aimodel"
	"github.com/testcoders/detestcoder/pkg/config/techstack"
	"os"
	"path"
)

var models = []string{"OpenAI"}
var modelVersions = []string{"gpt-4", "gpt-3.5-turbo"}
var yn = []string{"y", "n"}

type Config struct {
	AIModel   aimodel.AIModel     `mapstructure:"AIModel"`
	TechStack techstack.TechStack `mapstructure:"TechStack"`
}

type configPrompt struct {
	label    string
	errorMsg string
}

// ReadConfig loads the .detestcoder config file from the projects' working directory
func ReadConfig() (*Config, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(workingDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".detestcoder")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
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
	techStack := techstack.NewTechStack()

	workingDir, err := os.Getwd()
	cobra.CheckErr(err)

	getAiModelInput(aiModel)
	getTechstackInput(techStack, aiModel)

	// Configure Viper
	viper.AddConfigPath(workingDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".detestcoder")

	// Set config values
	viper.Set("AIModel", aiModel)
	viper.Set("TechStack", techStack)

	cfgFile := path.Join(workingDir, ".detestcoder.yaml")

	// Create the file if it does not exist.
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		_, err = os.Create(cfgFile)
		if err != nil {
			return err
		}
	}

	// Overwrite the existing config or Write a new one
	err = viper.WriteConfigAs(cfgFile)
	if err != nil {
		return err
	}

	return nil
}

// askUserToProceed checks whether the user wants to overwrite the existing setup
func askUserToProceed() string {
	proceedPrompt := configPrompt{
		label:    "It seems detestcoder is already initialized. You can update the existing setup this way. Do you wish to continue? ",
		errorMsg: "Select 'y' or 'n'.",
	}
	return getUserInputSelect(proceedPrompt, yn)
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
	modelNamePrompt := configPrompt{
		label:    "Which aimodel do you want to use? ",
		errorMsg: "Select an aimodel",
	}
	modelName := getUserInputSelect(modelNamePrompt, models)
	aiModel.SetModel(modelName)

	modelVersionPrompt := configPrompt{
		label:    fmt.Sprintf("Which version of '%v' would you like to use? ", modelName),
		errorMsg: "Provide a version",
	}
	modelVersion := getUserInputSelect(modelVersionPrompt, models)
	aiModel.SetModelVersion(modelVersion)

	apiKeyPrompt := configPrompt{
		label: fmt.Sprintf("Provide your '%v' API key ", modelName),
	}
	apiKey := getUserInputString(apiKeyPrompt, false, true)
	aiModel.SetApiKey(apiKey)
}

// getTechstackInput uses an interactive prompt to retrieve settings input from the user
func getTechstackInput(techstack *techstack.TechStack, aiModel *aimodel.AIModel) {
	automaticSetupPrompt := configPrompt{
		label:    fmt.Sprintf("Do you want to create the techstack automatically using '%v'? ", aiModel.AiModel),
		errorMsg: "Select 'y' or 'n'.",
	}
	var automaticSetup = getUserInputSelect(automaticSetupPrompt, yn)

	if automaticSetup == "n" {
		// Programming language
		programmingLanguagePrompt := configPrompt{
			label:    "Which programming language is your project written in? ",
			errorMsg: "Provide a programming language, like Java or Go",
		}
		programmingLanguage := getUserInputString(programmingLanguagePrompt, false, false)

		programmingLanguageVersionPrompt := configPrompt{
			label:    fmt.Sprintf("Which version of '%v' does your project use? ", programmingLanguage),
			errorMsg: "Provide a programming language version.",
		}
		programmingLanguageVersion := getUserInputString(programmingLanguageVersionPrompt, false, false)

		techstack.SetLanguage(programmingLanguage, programmingLanguageVersion)

		// Dependency manager
		dependencyManagerPrompt := configPrompt{
			label:    "Which dependency manager does your project use? ",
			errorMsg: "Provide a dependency manager, like Gradle or Maven",
		}
		dependencyManager := getUserInputString(dependencyManagerPrompt, false, false)

		dependencyManagerVersionPrompt := configPrompt{
			label:    fmt.Sprintf("Which version of '%v' does your project use? ", dependencyManager),
			errorMsg: "Provide a dependency manager version.",
		}
		dependencyManagerVersion := getUserInputString(dependencyManagerVersionPrompt, false, false)

		techstack.SetDependencyManager(dependencyManager, dependencyManagerVersion)

		// Framework
		frameworkPrompt := configPrompt{
			label:    "Which framework does your project use? ",
			errorMsg: "Provide a framework, like Spring Boot",
		}
		framework := getUserInputString(frameworkPrompt, false, false)

		frameworkVersionPrompt := configPrompt{
			label:    fmt.Sprintf("Which version of '%v' does your project use? ", framework),
			errorMsg: "Provide a framework version.",
		}
		frameworkVersion := getUserInputString(frameworkVersionPrompt, false, false)

		techstack.SetFramework(framework, frameworkVersion)

		// Test framework
		testFrameworkPrompt := configPrompt{
			label:    "Which test framework does your project use? ",
			errorMsg: "Provide a test framework, like Spring Boot",
		}
		testFramework := getUserInputString(testFrameworkPrompt, false, false)

		testFrameworkVersionPrompt := configPrompt{
			label:    fmt.Sprintf("Which version of '%v' does your project use? ", testFramework),
			errorMsg: "Provide a framework version.",
		}
		testFrameworkVersion := getUserInputString(testFrameworkVersionPrompt, false, false)

		techstack.SetTestFramework(testFramework, testFrameworkVersion)

		// Test dependencies
		addTestDependenciesPrompt := configPrompt{
			label:    "Do you want to add one or more test dependencies? ",
			errorMsg: "Select 'y' or 'n'.",
		}
		var addMoreTestDepencies = getUserInputSelect(addTestDependenciesPrompt, yn)

		for addMoreTestDepencies == "y" {
			testDependencyPrompt := configPrompt{
				label:    "Which test dependency your project use? ",
				errorMsg: "Provide a test dependency, like jUnit or testify.",
			}
			testDependency := getUserInputString(testDependencyPrompt, false, false)

			testDependencyVersionPrompt := configPrompt{
				label:    fmt.Sprintf("Which version of '%v' does your project use? ", testDependency),
				errorMsg: "Provide a test dependency version.",
			}
			testDependencyVersion := getUserInputString(testDependencyVersionPrompt, false, false)

			techstack.AddTestDependency(testDependency, testDependencyVersion)

			moreToAddPrompt := configPrompt{
				label:    "Do you want to add another test dependency? ",
				errorMsg: "Select 'y' or 'n'.",
			}
			addMoreTestDepencies = getUserInputSelect(moreToAddPrompt, yn)
		}
	} else {
		generateTechstackAutomatically(techstack, aiModel)
	}
}

func generateTechstackAutomatically(techstack *techstack.TechStack, aiModel *aimodel.AIModel) {
	fmt.Println("Not yet implemented.")
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

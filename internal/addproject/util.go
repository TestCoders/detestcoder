package addproject

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/testcoders/detestcoder/pkg/config"
	"github.com/testcoders/detestcoder/pkg/config/techstack"
	"github.com/testcoders/detestcoder/pkg/config/techstack/determine/programmingLanguage"
	"github.com/testcoders/detestcoder/pkg/config/techstack/determine/programmingLanguage/java/gradle"
	"github.com/testcoders/detestcoder/pkg/config/techstack/determine/programmingLanguage/java/maven"
	"github.com/testcoders/detestcoder/pkg/constants/project"
	"os"
	"path"
)

var yn = []string{"y", "n"}

// ReadConfig loads the .detestcoder initialize file from the projects' working directory
func ReadConfig() (*techstack.TechStack, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(workingDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".detestcoder.project")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var projectTechstack techstack.TechStack
	err = viper.Unmarshal(&projectTechstack)
	if err != nil {
		return nil, err
	}

	return &projectTechstack, nil
}

// WriteConfig writes the project settings to a .detestcoder.project.yaml file in the projects' working directory
// NOTE: any subsequent call to this will overwrite existing settings
func WriteConfig() error {
	if isDetestCoderProjectInitialized() {
		if askUserToProceed() == "n" {
			os.Exit(0)
		}
	}

	techStack := techstack.NewTechStack()

	workingDir, err := os.Getwd()
	cobra.CheckErr(err)

	getTechstackInput(techStack)

	// Configure Viper
	viper.AddConfigPath(workingDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".detestcoder.project")

	// Set initialize values
	viper.Set("language", techStack.Language)
	viper.Set("dependencymanager", techStack.DependencyManager)
	viper.Set("framework", techStack.Framework)
	viper.Set("testdependencies", techStack.TestDependencies)

	cfgFile := path.Join(workingDir, ".detestcoder.project.yaml")

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

// getTechstackInput uses an interactive prompt to retrieve settings input from the user
func getTechstackInput(ts *techstack.TechStack) {
	automaticSetupPrompt := config.ConfigPrompt{
		Label:    "Do you want to create the techstack automatically? ",
		ErrorMsg: "Select 'y' or 'n'.",
	}
	var automaticSetup = getUserInputSelect(automaticSetupPrompt, yn)

	if automaticSetup == "n" {
		generateTechstackManually(ts)
	} else {
		generateTechstackAutomatically(ts)
	}
}

func generateTechstackManually(ts *techstack.TechStack) {
	// Programming language
	programmingLanguagePrompt := config.ConfigPrompt{
		Label:    "Which programming language is your project written in? ",
		ErrorMsg: "Provide a programming language, like Java or Go",
	}
	projectProgrammingLanguage := getUserInputString(programmingLanguagePrompt, false, false)

	programmingLanguageVersionPrompt := config.ConfigPrompt{
		Label:    fmt.Sprintf("Which version of '%v' does your project use? ", projectProgrammingLanguage),
		ErrorMsg: "Provide a programming language version.",
	}
	programmingLanguageVersion := getUserInputString(programmingLanguageVersionPrompt, false, false)

	ts.SetLanguage(projectProgrammingLanguage, programmingLanguageVersion)

	// Dependency manager
	dependencyManagerPrompt := config.ConfigPrompt{
		Label:    "Which dependency manager does your project use? ",
		ErrorMsg: "Provide a dependency manager, like Gradle or Maven",
	}
	dependencyManager := getUserInputString(dependencyManagerPrompt, false, false)

	dependencyManagerVersionPrompt := config.ConfigPrompt{
		Label:    fmt.Sprintf("Which version of '%v' does your project use? ", dependencyManager),
		ErrorMsg: "Provide a dependency manager version.",
	}
	dependencyManagerVersion := getUserInputString(dependencyManagerVersionPrompt, false, false)

	ts.SetDependencyManager(dependencyManager, dependencyManagerVersion)

	// Framework
	frameworkPrompt := config.ConfigPrompt{
		Label:    "Which framework does your project use? ",
		ErrorMsg: "Provide a framework, like Spring Boot",
	}
	framework := getUserInputString(frameworkPrompt, false, false)

	frameworkVersionPrompt := config.ConfigPrompt{
		Label:    fmt.Sprintf("Which version of '%v' does your project use? ", framework),
		ErrorMsg: "Provide a framework version.",
	}
	frameworkVersion := getUserInputString(frameworkVersionPrompt, false, false)

	ts.SetFramework(framework, frameworkVersion)

	// Test dependencies
	addTestDependenciesPrompt := config.ConfigPrompt{
		Label:    "Do you want to add one or more test dependencies? ",
		ErrorMsg: "Select 'y' or 'n'.",
	}
	var addMoreTestDepencies = getUserInputSelect(addTestDependenciesPrompt, yn)

	for addMoreTestDepencies == "y" {
		testDependencyPrompt := config.ConfigPrompt{
			Label:    "Which test dependency your project use? ",
			ErrorMsg: "Provide a test dependency, like jUnit or testify.",
		}
		testDependency := getUserInputString(testDependencyPrompt, false, false)

		testDependencyVersionPrompt := config.ConfigPrompt{
			Label:    fmt.Sprintf("Which version of '%v' does your project use? ", testDependency),
			ErrorMsg: "Provide a test dependency version.",
		}
		testDependencyVersion := getUserInputString(testDependencyVersionPrompt, false, false)

		ts.AddTestDependency(testDependency, testDependencyVersion)

		moreToAddPrompt := config.ConfigPrompt{
			Label:    "Do you want to add another test dependency? ",
			ErrorMsg: "Select 'y' or 'n'.",
		}
		addMoreTestDepencies = getUserInputSelect(moreToAddPrompt, yn)
	}
}

func generateTechstackAutomatically(ts *techstack.TechStack) {
	projectDependencyManager := programmingLanguage.GetDependencyManager()
	projectProgrammingLanguage := programmingLanguage.GetProgrammingLanguages(projectDependencyManager)

	for _, language := range projectProgrammingLanguage {
		if language == project.JAVA || language == project.KOTLIN || language == project.SCALA {
			if projectDependencyManager == project.GRADLE {
				techStack := gradle.DetermineTechstack()
				*ts = *techStack
			} else if projectDependencyManager == project.MAVEN {
				techStack := maven.DetermineTechstack()
				*ts = *techStack
			} else {
				fmt.Println("This programming language is supported, but not this dependency manager.")
			}
		} else {
			fmt.Println("This programming language is not yet supported.")
		}
	}
}

// getUserInputSelect creates a prompt where the user can provide textual input.
func getUserInputString(cp config.ConfigPrompt, allowEmpty, mask bool) string {
	validate := func(input string) error {
		if !allowEmpty && len(input) <= 0 {
			return errors.New(cp.ErrorMsg)
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
		Label:     cp.Label,
		Templates: templates,
		Validate:  validate,
		Mask:      maskRune,
	}

	result, err := prompt.Run()
	cobra.CheckErr(err) // NOTE: use own check err?
	return result
}

// askUserToProceed checks whether the user wants to overwrite the existing setup
func askUserToProceed() string {
	proceedPrompt := config.ConfigPrompt{
		Label:    "It seems detestcoder is already initialized. You can update the existing setup this way. Do you wish to continue? ",
		ErrorMsg: "Select 'y' or 'n'.",
	}
	return config.GetUserInputSelect(proceedPrompt, yn)
}

// isDetestCoderProjectInitialized checks whether detestcoders is already initialized
func isDetestCoderProjectInitialized() bool {
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

// getUserInputSelect creates a prompt where the user can select any of the provided items
func getUserInputSelect(cp config.ConfigPrompt, items []string) string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label: cp.Label,
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

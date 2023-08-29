package config

import (
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type ConfigPrompt struct {
	Label    string
	ErrorMsg string
}

// GetUserInputString creates a prompt where the user can provide textual input.
func GetUserInputString(cp ConfigPrompt, allowEmpty, mask bool) string {
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

// GetUserInputSelect creates a prompt where the user can select any of the provided items
func GetUserInputSelect(cp ConfigPrompt, items []string) string {
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

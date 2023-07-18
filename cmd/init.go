/*
Copyright © 2023 TestCoders / DeTesters

*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/testcoders/detestcoder/internal/config"
)

// Struct to represent the API request payload
type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIRequest struct {
	Model     string          `json:"model"`
	Messages  []OpenAIMessage `json:"messages"`
	MaxTokens string          `json:"max_tokens"`
}

// Struct to represent the API response payload
type GPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your 'detestcoder' setup",
	Long:  "Use this command to generate a .detestcoder.yaml config file in your home directory.",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(config.WriteConfig())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

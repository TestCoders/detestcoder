package cmd

import (
	"github.com/spf13/cobra"
	"github.com/testcoders/detestcoder/internal/config"
)

var initPromptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Initialize your 'detestcoder' setup",
	Long:  "Use this command to generate a .detestcoder.yaml config file in your home directory.",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(config.PromptInit())
	},
}

func init() {
	rootCmd.AddCommand(initPromptCmd)
}

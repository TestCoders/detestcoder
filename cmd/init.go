/*
Copyright © 2023 TestCoders / DeTesters

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/testcoders/detestcoder/internal/app"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your 'detestcoder' setup",
	Long:  "Use this command to generate a .detestcoder.yaml config file in your home directory.",
	Run: func(cmd *cobra.Command, args []string) {
		app.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

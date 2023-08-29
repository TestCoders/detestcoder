/*
Copyright © 2023 TestCoders / DeTesters

*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/testcoders/detestcoder/internal/initialize"
)

// initCmd represents the initialize command
var initCmd = &cobra.Command{
	Use:   "initialize",
	Short: "Initialize your 'detestcoder' setup",
	Long:  "Use this command to files a .detestcoder.yaml initialize file in your home directory.",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(initialize.WriteConfig())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

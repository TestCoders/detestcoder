/*
Copyright © 2023 TestCoders, deTesters, TechChamps
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create initial config file with important details of your application.",
	Long: `Create initial config file with important details of your application.
The content of this file is based on your dependecy file.

It will output the following content:
- Programming language used in the application.
- Framework(s) utilized in the application.
- Unit test framework used, if any.
- Any other unit test plugins or dependencies mentioned in the file.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

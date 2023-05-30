package cmd

import (
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "testalyze [source-file] [test-file]",
	Short: "Analyze source and test files",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		sourceFile := args[0]
		testFile := ""
		if len(args) == 2 {
			testFile = args[1]
		}
		file.processFiles(sourceFile, testFile)
	},
}

func analyze() {
	rootCmd.AddCommand(testCmd)
}

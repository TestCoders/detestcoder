package cmd

import (
	"github.com/spf13/cobra"
	"github.com/testcoders/detestcoder/internal/file"
)

var analyzeCmd = &cobra.Command{
	Use:   "testalyze [source-file] [test-file]",
	Short: "Analyze source and test files",
	Long:  "Analyze the source file and optionally the test files. It writes back a file in the directory of the source file or, if specified, the test-file",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		sourceFile := args[0]
		testFile := ""
		if len(args) == 2 {
			testFile = args[1]
		}
		cobra.CheckErr(file.ProcessFiles(sourceFile, testFile))
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}

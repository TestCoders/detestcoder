package cmd

import (
	"github.com/spf13/cobra"
	"github.com/testcoders/detestcoder/internal/addproject"
)

var addprojectCmd = &cobra.Command{
	Use:   "addproject",
	Short: "Add a .detestcoder.project.yaml to your project root",
	Long:  `This command adds a configuration file to your project root.`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(addproject.WriteConfig())
	},
}

func init() {
	rootCmd.AddCommand(addprojectCmd)
}

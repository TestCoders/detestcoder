package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/testcoders/detestcoder/internal/addproject"
	"github.com/testcoders/detestcoder/internal/ai"
	"github.com/testcoders/detestcoder/internal/files"
	"github.com/testcoders/detestcoder/internal/initialize"
	"github.com/testcoders/detestcoder/pkg/config/techstack"
	"github.com/testcoders/detestcoder/pkg/constants/testType"
	"github.com/testcoders/detestcoder/pkg/promptbuilder"
)

var unitTest bool
var integrationTest bool
var e2eTest bool

func NewGenerateCmd() *cobra.Command {
	return generateCmd
}

var generateCmd = &cobra.Command{
	Use:   "generate [file] [context_of_file]",
	Short: "Generate tests for a given file with a given context",
	Long:  `This command generates unit, integration, or e2e tests for a given file with a given context.`,
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		aiModel, err := initialize.ReadConfig()
		if err != nil {
			fmt.Printf("Failed to read initialize: %v", err)
			return
		}

		techStack, err := addproject.ReadConfig()
		if err != nil {
			fmt.Printf("Failed to read initialize: %v", err)
			return
		}

		pb := promptbuilder.NewPromptBuilder()

		var codeSnippetContext string
		if len(args) > 1 {
			codeSnippetContext = args[1]
		}
		pb.AddCodeSnippetContext(codeSnippetContext)

		if cmd.Flags().NFlag() == 0 {
			fmt.Printf("No flags provided, defaulting to unit test for file: %s\n", file)
			pb.AddKindOfTest(testType.UT)
		} else {
			if unitTest {
				fmt.Printf("Generating unit tests for file: %s\n", file)
				pb.AddKindOfTest(testType.UT)
			}
			if integrationTest {
				fmt.Printf("Generating integration tests for file: %s\n", file)
				pb.AddKindOfTest(testType.IT)
			}
			if e2eTest {
				fmt.Printf("Generating e2e tests for file: %s\n", file)
				pb.AddKindOfTest(testType.E2E)
			}
		}

		files.ReadContentsOfFileAndAddCodeSnippet(pb, file)

		pb.AddProgrammingLanguage(techStack.Language.Name)
		pb.AddProgrammingLanguageVersion(techStack.Language.Version)
		pb.AddDependencyManager(techStack.DependencyManager.Name + " " + techStack.DependencyManager.Version)
		pb.AddFrameworks(techStack.Framework.Name + " " + techStack.Framework.Version)
		pb.AddTestFramework(techStack.TestFramework.Name + " " + techStack.TestFramework.Version)
		pb.AddTestDependencies(techstack.GetTestDependencies(*techStack))

		response, err := ai.SendPrompt(ai.NewService().GetService(pb.Build(), *aiModel), *aiModel)
		if err != nil {
			panic(err)
		}

		fmt.Println(response.Content)
		files.WriteOutputToFile(*response, file)
	},
}

func init() {
	generateCmd.Flags().BoolVarP(&unitTest, "unittest", "u", false, "Generate unit tests")
	generateCmd.Flags().BoolVarP(&integrationTest, "integrationtest", "i", false, "Generate integration tests")
	generateCmd.Flags().BoolVarP(&e2eTest, "e2etest", "e", false, "Generate e2e tests")

	rootCmd.AddCommand(generateCmd)
}

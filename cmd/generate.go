package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/testcoders/detestcoder/pkg/constants/testType"
	"github.com/testcoders/detestcoder/pkg/promptbuilder"
	"github.com/testcoders/detestcoder/pkg/techstack"
	"log"
	"os"
	"strings"
	"time"
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

		ts := techstack.GetCurrentTechStack()

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

		readContentsOfFileAndAddCodeSnippet(pb, file)

		pb.AddProgrammingLanguage(ts.TechStack.Language.Name)
		pb.AddProgrammingLanguageVersion(ts.TechStack.Language.Version)
		pb.AddDependencyManager(ts.TechStack.DependencyManager.Name + " " + ts.TechStack.DependencyManager.Version)
		pb.AddFrameworks(ts.TechStack.Framework.Name + " " + ts.TechStack.Framework.Version)
		pb.AddTestFramework(ts.TechStack.TestFramework.Name + " " + ts.TechStack.TestFramework.Version)
		pb.AddTestDependencies(getTestDependencies(ts))

		prompt := pb.Build()

		// TODO: actually send this to OpenAI and do the magic from there...
		writeToFile(prompt)
		fmt.Printf(prompt)
	},
}

func writeToFile(data string) {
	timestamp := time.Now().Unix()

	// create the directory if it doesn't exist
	dir := "generatedPrompts/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	filename := fmt.Sprintf("%sprompt_%d", dir, timestamp)

	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

func readContentsOfFileAndAddCodeSnippet(pb *promptbuilder.PromptBuilder, filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	code := string(bytes)
	pb.AddCodeSnippet(code)
}

func getTestDependencies(ts *techstack.TechStack) string {
	dependencies := make([]string, len(ts.TechStack.TestDependencies))

	for i, dependency := range ts.TechStack.TestDependencies {
		dependencies[i] = fmt.Sprintf("%s %s", dependency.Name, dependency.Version)
	}

	dependenciesString := strings.Join(dependencies, ", ")

	return dependenciesString
}

func init() {
	generateCmd.Flags().BoolVarP(&unitTest, "unittest", "u", false, "Generate unit tests")
	generateCmd.Flags().BoolVarP(&integrationTest, "integrationtest", "i", false, "Generate integration tests")
	generateCmd.Flags().BoolVarP(&e2eTest, "e2etest", "e", false, "Generate e2e tests")

	rootCmd.AddCommand(generateCmd)
}

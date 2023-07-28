package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/testcoders/detestcoder/pkg/constants/testType"
	"github.com/testcoders/detestcoder/pkg/promptbuilder"
	"github.com/testcoders/detestcoder/pkg/techstack"
	"strings"
)

var (
	unitTest        bool
	integrationTest bool
	e2eTest         bool
	generateCmd     = defineGenerateCmd()
)

func NewGenerateCmd() *cobra.Command {
	return generateCmd
}

func defineGenerateCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "generate [file] [context_of_file]",
		Short: "Generate tests for a given file with a given context",
		Long:  `This command generates unit, integration, or e2e tests for a given file with a given context.`,
		Args:  cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			ts := techstack.GetCurrentTechStack()
			pb := promptbuilder.NewPromptBuilder()
			pb.AddCodeSnippet(file) // TODO: func for readValueOfFile

			var codeSnippetContext string
			if len(args) > 1 {
				codeSnippetContext = args[1]
			}
			pb.AddCodeSnippetContext(codeSnippetContext)
			addTestTypePB(pb, cmd, file)
			specifyTechStacksPB(pb, ts)
			prompt := pb.Build()
			fmt.Printf(prompt)
		},
	}

	// Setting up flags
	cmd.Flags().BoolVarP(&unitTest, "unittest", "u", false, "Generate unit tests")
	cmd.Flags().BoolVarP(&integrationTest, "integrationtest", "i", false, "Generate integration tests")
	cmd.Flags().BoolVarP(&e2eTest, "e2etest", "e", false, "Generate e2e tests")

	return cmd
}

func addTestTypePB(pb *promptbuilder.PromptBuilder, cmd *cobra.Command, file string) {
	if cmd.Flags().NFlag() == 0 {
		fmt.Printf("No flags provided, defaulting to unit test for file: %s\n", file)
		pb.AddKindOfTest(testType.UT)
		return
	}

	testTypeMapping := map[bool]struct {
		full   string
		abbrev string
	}{
		unitTest:        {full: "unit tests", abbrev: testType.UT},
		integrationTest: {full: "integration tests", abbrev: testType.IT},
		e2eTest:         {full: "end-to-end tests", abbrev: testType.E2E},
	}

	for test, kind := range testTypeMapping {
		if test {
			fmt.Printf("Generating %s for file: %s\n", kind.full, file)
			pb.AddKindOfTest(kind.abbrev)
		}
	}
}

func specifyTechStacksPB(pb *promptbuilder.PromptBuilder, ts *techstack.TechStack) {
	pb.AddProgrammingLanguage(ts.TechStack.Language.Name)
	pb.AddProgrammingLanguageVersion(ts.TechStack.Language.Version)
	pb.AddDependencyManager(ts.TechStack.DependencyManager.Name + " " + ts.TechStack.DependencyManager.Version)
	pb.AddFrameworks(ts.TechStack.Framework.Name + " " + ts.TechStack.Framework.Version)
	pb.AddTestFramework(ts.TechStack.TestFramework.Name + " " + ts.TechStack.TestFramework.Version)
	pb.AddTestDependencies(getTestDependencies(ts))
}

func getTestDependencies(ts *techstack.TechStack) string {
	dependencies := make([]string, len(ts.TechStack.TestDependencies))

	for i, dependency := range ts.TechStack.TestDependencies {
		dependencies[i] = fmt.Sprintf("%s %s", dependency.Name, dependency.Version)
	}

	dependenciesString := strings.Join(dependencies, ", ")

	return dependenciesString
}

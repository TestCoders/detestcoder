package promptbuilder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcoders/detestcoder/pkg/constants"
	"testing"
)

func TestNewPromptBuilder(t *testing.T) {
	pb := NewPromptBuilder()
	assert.NotNil(t, pb)
	assert.Equal(t, basePrompt, pb.basePrompt)
	assert.NotNil(t, pb.variables)
}

func TestAddProgrammingLanguage(t *testing.T) {
	pb := NewPromptBuilder()
	pb.AddProgrammingLanguage("Go")
	assert.Equal(t, "Go", pb.variables[constants.ProgrammingLanguage])
}

func TestAddProgrammingLanguageVersion(t *testing.T) {
	pb := NewPromptBuilder()
	pb.AddProgrammingLanguageVersion("1.20")
	assert.Equal(t, "1.20", pb.variables[constants.ProgrammingLanguageVersion])
}

func TestAddFrameworks(t *testing.T) {
	pb := NewPromptBuilder()
	pb.AddFrameworks([]string{"cobra"})
	assert.Equal(t, "cobra", pb.variables[constants.Frameworks])
}

func TestAddUnitTestFramework(t *testing.T) {
	pb := NewPromptBuilder()
	pb.AddUnitTestFramework("testify")
	assert.Equal(t, "testify", pb.variables[constants.TestFramework])
}

func TestAddUnitTestDependencies(t *testing.T) {
	pb := NewPromptBuilder()
	pb.AddUnitTestDependencies([]string{"testify"})
	assert.Equal(t, "testify", pb.variables[constants.TestDependencies])
}

func TestAddCodeSnippet(t *testing.T) {
	pb := NewPromptBuilder()
	pb.AddCodeSnippet("func testFunc() {}")
	assert.Equal(t, "func testFunc() {}", pb.variables[constants.CodeSnippet])
}

func TestAddCodeSnippetContext(t *testing.T) {
	pb := NewPromptBuilder()
	pb.AddCodeSnippetContext("A test function")
	assert.Equal(t, "A test function", pb.variables[constants.CodeSnippetContext])
}

func TestAddKindOfTest(t *testing.T) {
	pb := NewPromptBuilder()
	pb.AddKindOfTest("Unit test")
	assert.Equal(t, "Unit test", pb.variables[constants.KindOfTest])
}

func TestBuild(t *testing.T) {
	pb := NewPromptBuilder()
	pb.AddProgrammingLanguage("Go")
	pb.AddProgrammingLanguageVersion("1.20")
	pb.AddFrameworks([]string{"cobra"})
	pb.AddUnitTestFramework("testify")
	pb.AddUnitTestDependencies([]string{"testify"})
	pb.AddCodeSnippet("func testFunc() {}")
	pb.AddCodeSnippetContext("A test function")
	pb.AddKindOfTest("Unit test")

	result := pb.Build()

	fmt.Println(result)

	assert.Contains(t, result, "Go")
	assert.Contains(t, result, "1.20")
	assert.Contains(t, result, "cobra")
	assert.Contains(t, result, "testify")
	assert.Contains(t, result, "func testFunc() {}")
	assert.Contains(t, result, "A test function")
	assert.Contains(t, result, "Unit test")
}

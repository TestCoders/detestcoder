package promptbuilder

import (
	"fmt"
	"github.com/testcoders/detestcoder/pkg/constants"
	"strings"
)

const basePrompt = "" +
	"You are an experienced test automation engineer. I need you to help me write {" + constants.KindOfTest + "}s for the following code snippet from my project:\n\n" +
	"{" + constants.CodeSnippet + "}\n\n" +
	"This code has the following context: {" + constants.CodeSnippetContext + "}\n\n" +
	"The code is written in {" + constants.ProgrammingLanguage + "} with version {" + constants.ProgrammingLanguageVersion + "}.\n\n" +
	"It's built using the following frameworks: {" + constants.Frameworks + "}\n\n" +
	"It uses the following test frameworks \"{" + constants.TestFramework + "}\" and dependencies \"{" + constants.TestDependencies + "}\"\n\n" +
	"Please provide only the test code without ``` without any explaining around it. \n\n" +
	"Explain how you got to the test case using comments above each test function."

type PromptBuilder struct {
	basePrompt string
	variables  map[string]string
}

func NewPromptBuilder() *PromptBuilder {
	return &PromptBuilder{
		basePrompt: basePrompt,
		variables:  make(map[string]string),
	}
}

func (p *PromptBuilder) AddProgrammingLanguage(value string) {
	p.addVariable(constants.ProgrammingLanguage, value)
}

func (p *PromptBuilder) AddProgrammingLanguageVersion(value string) {
	p.addVariable(constants.ProgrammingLanguageVersion, value)
}

func (p *PromptBuilder) AddFrameworks(values []string) {
	commaSeparatedValues := strings.Join(values, ", ")
	p.addVariable(constants.Frameworks, commaSeparatedValues)
}

func (p *PromptBuilder) AddUnitTestFramework(value string) {
	p.addVariable(constants.TestFramework, value)
}

func (p *PromptBuilder) AddUnitTestDependencies(values []string) {
	commaSeparatedValues := strings.Join(values, ", ")
	p.addVariable(constants.TestDependencies, commaSeparatedValues)
}

func (p *PromptBuilder) AddCodeSnippet(value string) {
	p.addVariable(constants.CodeSnippet, value)
}

func (p *PromptBuilder) AddCodeSnippetContext(value string) {
	p.addVariable(constants.CodeSnippetContext, value)
}

func (p *PromptBuilder) AddKindOfTest(value string) {
	p.addVariable(constants.KindOfTest, value)
}

func (p *PromptBuilder) addVariable(key, value string) {
	p.variables[key] = value
}

func (p *PromptBuilder) Build() string {
	result := p.basePrompt

	for key, value := range p.variables {
		placeholder := fmt.Sprintf("{%s}", key)
		result = strings.Replace(result, placeholder, value, -1)
	}

	return result
}

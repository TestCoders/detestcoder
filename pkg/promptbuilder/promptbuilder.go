package promptbuilder

import (
	"fmt"
	"github.com/testcoders/detestcoder/pkg/constants/promptConstants"
	"strings"
)

const basePrompt = "" +
	"You are an experienced test automation engineer. I need you to help me write {" + promptConstants.KindOfTest + "}s for the following code snippet from my project:\n\n" +
	"{" + promptConstants.CodeSnippet + "}\n\n" +
	"This code has the following context: {" + promptConstants.CodeSnippetContext + "}\n\n" +
	"The code is written in {" + promptConstants.ProgrammingLanguage + "} with version {" + promptConstants.ProgrammingLanguageVersion + "}.\n\n" +
	"It uses the following dependency manager (ignore this when empty): {" + promptConstants.DependencyManager + "}.\n\n" +
	"It's built using the following frameworks: {" + promptConstants.Frameworks + "}\n\n" +
	"It uses the following test frameworks \"{" + promptConstants.TestFramework + "}\" and dependencies \"{" + promptConstants.TestDependencies + "}\"\n\n" +
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
	p.addVariable(promptConstants.ProgrammingLanguage, value)
}

func (p *PromptBuilder) AddProgrammingLanguageVersion(value string) {
	p.addVariable(promptConstants.ProgrammingLanguageVersion, value)
}

func (p *PromptBuilder) AddDependencyManager(value string) {
	p.addVariable(promptConstants.DependencyManager, value)
}

func (p *PromptBuilder) AddFrameworks(value string) {
	p.addVariable(promptConstants.Frameworks, value)
}

func (p *PromptBuilder) AddTestFramework(value string) {
	p.addVariable(promptConstants.TestFramework, value)
}

func (p *PromptBuilder) AddTestDependencies(value string) {
	p.addVariable(promptConstants.TestDependencies, value)
}

func (p *PromptBuilder) AddCodeSnippet(value string) {
	p.addVariable(promptConstants.CodeSnippet, value)
}

func (p *PromptBuilder) AddCodeSnippetContext(value string) {
	p.addVariable(promptConstants.CodeSnippetContext, value)
}

func (p *PromptBuilder) AddKindOfTest(value string) {
	p.addVariable(promptConstants.KindOfTest, value)
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

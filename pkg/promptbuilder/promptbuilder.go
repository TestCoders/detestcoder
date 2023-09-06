package promptbuilder

import (
	"github.com/testcoders/detestcoder/pkg/constants/promptConstants"
	"strings"
)

const basePrompt = "" +
	"You are an experienced test automation engineer. I need you to help me write {" + promptConstants.TestType + "}s for the following {" + promptConstants.ProgrammingLanguage + "} code snippet from my project:\n\n" +
	"{" + promptConstants.CodeSnippet + "}\n\n" +
	"This code has the following context: {" + promptConstants.CodeSnippetContext + "}\n\n" +
	"The code is written in {" + promptConstants.ProgrammingLanguage + "} with version {" + promptConstants.ProgrammingLanguageVersion + "}.\n\n" +
	"It uses the following dependency manager (ignore this when empty): {" + promptConstants.DependencyManager + "}.\n\n" +
	"It's built using the following frameworks: {" + promptConstants.Frameworks + "}\n\n" +
	"It uses the following test dependencies \"{" + promptConstants.TestDependencies + "}\". Please make sure all imports are set correctly.\n\n" +
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
	p.addVariable(promptConstants.TestType, value)
}

func (p *PromptBuilder) addVariable(key, value string) {
	p.variables[key] = value
}

func (p *PromptBuilder) Build() string {
	var result strings.Builder

	last := 0
	for start := 0; start < len(p.basePrompt); start++ {
		if p.basePrompt[start] == '{' {
			end := strings.IndexByte(p.basePrompt[start:], '}')
			if end < 0 {
				break
			}
			end += start

			key := p.basePrompt[start+1 : end]
			value, ok := p.variables[key]
			if ok {
				result.WriteString(p.basePrompt[last:start])
				result.WriteString(value)
				last = end + 1
			}
		}
	}
	result.WriteString(p.basePrompt[last:])

	return result.String()
}

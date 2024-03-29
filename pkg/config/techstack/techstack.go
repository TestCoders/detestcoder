package techstack

import (
	"fmt"
	"strings"
)

type TechStack struct {
	Language struct {
		Name            string `yaml:"name"`
		Version         string `yaml:"version"`
		Compiler        string `yaml:"compiler"`
		CompilerVersion string `yaml:"compilerversion"`
	} `yaml:"language"`
	DependencyManager struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"dependencymanager"`
	Framework struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"framework"`
	TestDependencies []struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"testdependencies"`
}

func NewTechStack() *TechStack {
	return &TechStack{
		Language: struct {
			Name            string `yaml:"name"`
			Version         string `yaml:"version"`
			Compiler        string `yaml:"compiler"`
			CompilerVersion string `yaml:"compilerversion"`
		}{},
		DependencyManager: struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		}{},
		Framework: struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		}{},
		TestDependencies: []struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		}{},
	}
}

func (t *TechStack) SetDependencyManager(name string, version string) {
	t.DependencyManager.Name = name
	t.DependencyManager.Version = version
}

func (t *TechStack) SetLanguage(name string, version string) {
	t.Language.Name = name
	t.Language.Version = version
}

func (t *TechStack) SetCompiler(name string, version string) {
	t.Language.Compiler = name
	t.Language.CompilerVersion = version
}

func (t *TechStack) SetFramework(name string, version string) {
	t.Framework.Name = name
	t.Framework.Version = version
}

func (t *TechStack) AddTestDependency(name string, version string) {
	unitTestDependency := struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}{
		Name:    name,
		Version: version,
	}
	t.TestDependencies = append(t.TestDependencies, unitTestDependency)
}

func GetTestDependencies(ts TechStack) string {
	dependencies := make([]string, len(ts.TestDependencies))

	for i, dependency := range ts.TestDependencies {
		dependencies[i] = fmt.Sprintf("%s %s", dependency.Name, dependency.Version)
	}

	dependenciesString := strings.Join(dependencies, ", ")

	return dependenciesString
}

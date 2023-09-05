package maven

import (
	"fmt"
	"github.com/testcoders/detestcoder/pkg/config/techstack"
	"github.com/testcoders/detestcoder/pkg/constants/project"
	"github.com/testcoders/detestcoder/pkg/constants/project/java"
	"github.com/vifraa/gopom"
	"log"
	"strings"
)

func DetermineTechstack() *techstack.TechStack {
	parsedPom, err := gopom.Parse("pom.xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parsedPom)

	ts := techstack.NewTechStack()

	// TODO: Check if there is a parent pom or child pom(s)
	dependencies := getDependencies(*parsedPom)

	pl, plv, c, cv := determineLanguageAndVersion(parsedPom)
	ts.SetLanguage(pl, plv)
	if c != "" {
		ts.SetCompiler(c, cv)
	}
	ts.SetDependencyManager(project.MAVEN, "")
	ts.SetFramework(determineFramework(dependencies))

	testDependencies := getTestDependencies(dependencies)

	for _, testDependency := range testDependencies {
		ts.AddTestDependency(getTestDependency(testDependency))
	}

	return ts
}

func getDependencies(pom gopom.Project) *[]gopom.Dependency {
	return pom.Dependencies
}

func determineLanguageAndVersion(pom *gopom.Project) (string, string, string, string) {
	lang := ""
	langVersion := ""
	compiler := ""
	compilerVersion := ""

	propMap := make(map[string]string)
	for key, entry := range pom.Properties.Entries {
		propMap[key] = entry
	}

	if _, ok := propMap["kotlin.version"]; ok {
		lang = project.KOTLIN
		langVersion = propMap["kotlin.version"]
		compiler = project.JAVA
		compilerVersion = propMap["java.version"]
	} else if _, ok := propMap["scala.version"]; ok {
		lang = project.SCALA
		langVersion = propMap["scala.version"]
		compiler = project.JAVA
		compilerVersion = propMap["java.version"]
	} else if _, ok := propMap["java.version"]; ok {
		lang = project.JAVA
		langVersion = propMap["java.version"]
	}

	return lang, langVersion, compiler, compilerVersion
}

func determineFramework(dependencies *[]gopom.Dependency) (string, string) {
	framework := ""
	frameworkVersion := ""

	for _, fw := range java.Frameworks {
		for _, dep := range *dependencies {
			if strings.Contains(*dep.ArtifactID, fw) {
				framework = *dep.ArtifactID
				frameworkVersion = *dep.Version
			}
		}

		// Break outer loop if framework found
		if framework != "" && frameworkVersion != "" {
			break
		}
	}

	return framework, frameworkVersion
}

func getTestDependencies(dependencies *[]gopom.Dependency) []gopom.Dependency {
	var testDeps []gopom.Dependency

	for _, dep := range *dependencies {
		if dep.Scope != nil && strings.Contains(*dep.Scope, "test") {
			testDeps = append(testDeps, dep)
		}
	}

	return testDeps
}

func getTestDependency(dependency gopom.Dependency) (string, string) {
	version := ""

	if dependency.Version != nil && *dependency.Version != "" {
		version = *dependency.Version
	}

	return *dependency.ArtifactID, version
}

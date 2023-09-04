package gradle

import (
	"github.com/testcoders/detestcoder/pkg/config/techstack"
	"github.com/testcoders/detestcoder/pkg/constants/project"
	"github.com/testcoders/detestcoder/pkg/constants/project/java"
	"log"
	"os"
	"regexp"
	"strings"
)

func DetermineTechstack() *techstack.TechStack {
	content, err := os.ReadFile("build.gradle") // Read the file in the current directory
	if err != nil {
		log.Fatal(err)
	}

	ts := techstack.NewTechStack()

	dependencies := getDependencies(string(content))

	ts.SetLanguage(determineLanguageAndVersion(string(content)))
	ts.SetDependencyManager(project.GRADLE, "")
	ts.SetFramework(determineFramework(dependencies))

	testDependencies := getTestDependencies(dependencies)
	for _, testDependency := range testDependencies {
		ts.AddTestDependency(getTestDepency(string(content), testDependency))
	}

	return ts
}

func getDependencies(dependencyManagerFile string) []string {
	re := regexp.MustCompile(`(compileOnly|annotationProcessor|testImplementation|testRuntimeOnly|implementation).+(\(.*?\)|'.*?'|".*?")`)
	matches := re.FindAllString(dependencyManagerFile, -1)
	return matches
}

func determineLanguageAndVersion(dependencyManagerFile string) (string, string) {
	lang := ""
	langVersion := ""

	javaPluginRegexp := regexp.MustCompile(`id 'java'`)
	kotlinPluginRegexp := regexp.MustCompile(`id 'kotlin'`)
	scalaPluginRegexp := regexp.MustCompile(`id 'scala'`)
	if javaPluginRegexp.MatchString(dependencyManagerFile) {
		lang = project.JAVA
		javaVersionRegexp := regexp.MustCompile(`sourceCompatibility = '(\d*)'`)
		javaVersionMatch := javaVersionRegexp.FindStringSubmatch(dependencyManagerFile)
		if len(javaVersionMatch) > 1 {
			langVersion = javaVersionMatch[1]
		}
	}
	if kotlinPluginRegexp.MatchString(dependencyManagerFile) {
		lang = project.KOTLIN
		javaVersionRegexp := regexp.MustCompile(`sourceCompatibility = '(\d*)'`)
		javaVersionMatch := javaVersionRegexp.FindStringSubmatch(dependencyManagerFile)
		if len(javaVersionMatch) > 1 {
			langVersion = javaVersionMatch[1]
		}
	}
	if scalaPluginRegexp.MatchString(dependencyManagerFile) {
		lang = project.SCALA
		javaVersionRegexp := regexp.MustCompile(`sourceCompatibility = '(\d*)'`)
		javaVersionMatch := javaVersionRegexp.FindStringSubmatch(dependencyManagerFile)
		if len(javaVersionMatch) > 1 {
			langVersion = javaVersionMatch[1]
		}
	}
	return lang, langVersion
}

func determineFramework(dependencies []string) (string, string) {
	framework := ""
	frameworkVersion := ""

	for _, fw := range java.Frameworks {
		for _, dep := range dependencies {
			if strings.Contains(dep, fw) {
				// Trim dependency function and parentheses or quotes
				dep = strings.Trim(dep[strings.Index(dep, "(")+1:len(dep)-1], "' ")

				// Split the depedency
				splitDep := strings.Split(dep, ":")

				if len(splitDep) > 2 {
					framework = splitDep[1]
					frameworkVersion = splitDep[2]
					break
				}
			}
		}

		// Break outer loop if framework found
		if framework != "" && frameworkVersion != "" {
			break
		}
	}

	return framework, frameworkVersion
}

func getTestDependencies(dependencies []string) []string {
	var testDeps []string
	for _, dep := range dependencies {
		if strings.HasPrefix(dep, "testImplementation") || strings.HasPrefix(dep, "testRuntimeOnly") {
			testDeps = append(testDeps, dep)
		}
	}
	return testDeps
}

func getTestDepency(gradleBuildFile string, testDependency string) (string, string) {
	testDep := ""
	testDepVersion := ""

	mapRe := regexp.MustCompile(`group:\s*'([^']*)',\s*name:\s*'([^']*)',\s*version:\s*'([^']*)'|["']([^"']*)["']`)

	mapMatch := mapRe.FindStringSubmatch(testDependency)

	if mapMatch[1] != "" || mapMatch[2] != "" && mapMatch[3] != "" {
		testDep = mapMatch[2]
		testDepVersion = mapMatch[3]
	} else {
		stringRe := regexp.MustCompile(`["']([^"']*)["']`)
		stringMatch := stringRe.FindStringSubmatch(testDependency)
		splitDep := strings.Split(stringMatch[1], ":")
		if len(splitDep) > 2 {
			testDep = splitDep[1]
			testDepVersion = splitDep[2]
		}
	}

	if checkIfVariableIsAbstract(testDepVersion) {
		testDepVersion = strings.ReplaceAll(testDepVersion, "${", "")
		testDepVersion = strings.ReplaceAll(testDepVersion, "}", "")
		testDepVersion = findActualTestDepVersionInBuildFile(gradleBuildFile, testDepVersion)
	}

	return testDep, testDepVersion
}

func checkIfVariableIsAbstract(input string) bool {
	return strings.HasPrefix(input, "${") && strings.HasSuffix(input, "}")
}

func findActualTestDepVersionInBuildFile(gradleBuildFile string, testDepVersion string) string {
	re := regexp.MustCompile(testDepVersion + `\s*=\s*'([^']*)'`)

	match := re.FindStringSubmatch(gradleBuildFile)

	return match[1]
}

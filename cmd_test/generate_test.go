package cmd_test

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/testcoders/detestcoder/cmd"
	"github.com/testcoders/detestcoder/pkg/config/aimodel"
	"github.com/testcoders/detestcoder/pkg/config/techstack"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"testing"
)

func setup() {
	createTempTestfile("testfile.java")

	am := aimodel.NewAiModel()
	am.SetModel("Mock")
	am.SetModelVersion("4")
	am.SetApiKey("12345678")

	createTempDetestcoderYaml(*am)

	ts := techstack.NewTechStack()
	ts.SetDependencyManager("Maven", "3")
	ts.SetLanguage("Java", "20")
	ts.SetFramework("Spring Boot", "3.0")
	ts.AddTestDependency("AssertJ", "3.1.1")
	ts.AddTestDependency("Spring Boot Test", "3")

	createTempDetestcoderProjectYaml(*ts)
}

func createTempTestfile(testfile string) {
	content := []byte(`
		Iterator<Map<String, Object>> feeder =
 			Stream.files((Supplier<Map<String, Object>>) () -> {
     		String email = RandomStringUtils.randomAlphanumeric(20) + "@foo.com";
     		return Collections.singletonMap("email", email);
			}
		).iterator();`)
	err := os.WriteFile(testfile, content, 0644)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
}

func createTempDetestcoderProjectYaml(ts techstack.TechStack) {
	// Marshal the ts struct into YAML
	bytes, err := yaml.Marshal(ts)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Write the YAML to the .detestcoder file
	err = os.WriteFile(".detestcoder.project.yaml", bytes, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func createTempDetestcoderYaml(am aimodel.AIModel) {
	workingDir, err := os.Getwd()
	cobra.CheckErr(err)

	err = os.Setenv("HOME", workingDir)
	if err != nil {
		return
	}

	// Marshal the ts struct into YAML
	bytes, err := yaml.Marshal(am)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Write the YAML to the .detestcoder file
	err = os.WriteFile(".detestcoder.yaml", bytes, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func teardown() {
	err := os.Remove(".detestcoder.yaml")
	if err != nil {
		log.Fatalf("Failed to remove .detestcoder.yaml: %v", err)
	}
	err = os.Unsetenv("HOME")
	if err != nil {
		return
	}
	err = os.Remove(".detestcoder.project.yaml")
	if err != nil {
		log.Fatalf("Failed to remove .detestcoder.project.yaml: %v", err)
	}
	err = os.Remove("testfile.java")
	if err != nil {
		log.Fatalf("Failed to remove testfile.java: %v", err)
	}
	err = os.RemoveAll("generatedOutput")
	if err != nil {
		log.Fatalf("Failed to remove the generatedOutput directory: %v", err)
	}
	err = os.RemoveAll("generatedPrompts")
	if err != nil {
		log.Fatalf("Failed to remove the generatedPrompts directory: %v", err)
	}
}

func getGeneratedOutputFromDir() string {
	output := ""

	workingDir, err := os.Getwd()
	cobra.CheckErr(err)

	files, err := os.ReadDir(workingDir + "/generatedOutput/")
	if err != nil {
		fmt.Println(err)
		return output
	}

	for _, file := range files {
		if !file.IsDir() {
			content, err := os.ReadFile(workingDir + "/generatedOutput/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			output = string(content)
		}
	}

	return output
}

// Test case for default scenario where no flags are provided. The system should default to unit tests
func TestGenerateCmd_NoFlags(t *testing.T) {
	setup()
	defer teardown()

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.NewGenerateCmd())

	// Set the file argument
	rootCmd.SetArgs([]string{"generate", "testfile.java", "A file to test the test"})
	err := rootCmd.Execute()
	assert.NoError(t, err)

	output := getGeneratedOutputFromDir()

	assert.Contains(t, output, "```java\nimport org.apache.commons.lang3.RandomStringUtils;")
}

// Test case for when only the unit test flag is set
func TestGenerateCmd_UnitTestFlag(t *testing.T) {
	setup()
	defer teardown()

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.NewGenerateCmd())

	// Set the file argument
	rootCmd.SetArgs([]string{"generate", "testfile.java", "--unittest"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	output := getGeneratedOutputFromDir()

	assert.Contains(t, output, "```java\nimport org.apache.commons.lang3.RandomStringUtils;")
}

// Test case for when only the integration test flag is set
func TestGenerateCmd_IntegrationTestFlag(t *testing.T) {
	setup()
	defer teardown()

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.NewGenerateCmd())

	// Set the file argument
	rootCmd.SetArgs([]string{"generate", "testfile.java", "--integrationtest"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	output := getGeneratedOutputFromDir()

	assert.Contains(t, output, "```java\nimport org.apache.commons.lang3.RandomStringUtils;")
}

// Test case for when only the e2e test flag is set
func TestGenerateCmd_E2ETestFlag(t *testing.T) {
	setup()
	defer teardown()

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.NewGenerateCmd())

	// Set the file argument
	rootCmd.SetArgs([]string{"generate", "testfile.java", "--e2etest"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	output := getGeneratedOutputFromDir()

	assert.Contains(t, output, "```java\nimport org.apache.commons.lang3.RandomStringUtils;")
}

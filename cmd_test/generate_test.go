package cmd_test

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/testcoders/detestcoder/cmd"
	"github.com/testcoders/detestcoder/internal/config"
	"github.com/testcoders/detestcoder/pkg/config/aimodel"
	"github.com/testcoders/detestcoder/pkg/config/techstack"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"testing"
)

func setup() *techstack.TechStack {
	createTempTestfile("testfile.java")

	ts := techstack.NewTechStack()
	ts.SetDependencyManager("Maven", "3")
	ts.SetLanguage("Java", "20")
	ts.SetFramework("Spring Boot", "3.0")
	ts.SetTestFramework("jUnit", "5")
	ts.AddTestDependency("AssertJ", "3.1.1")
	ts.AddTestDependency("Spring Boot Test", "3")

	am := aimodel.NewAiModel()
	am.SetModel("Mock")
	am.SetModelVersion("4")
	am.SetApiKey("12345678")

	myConfig := config.Config{
		AIModel:   *am,
		TechStack: *ts,
	}

	createTempDetestcoderYaml(myConfig)

	return ts
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

func createTempDetestcoderYaml(config config.Config) {
	// Marshal the ts struct into YAML
	bytes, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Write the YAML to the .detestcoder file
	err = os.WriteFile(".detestcoder", bytes, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func teardown() {
	err := os.Remove(".detestcoder")
	if err != nil {
		log.Fatalf("Failed to remove .detestcoder: %v", err)
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

func captureOutput(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		buf.ReadFrom(r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	return out
}

// Test case for default scenario where no flags are provided. The system should default to unit tests
func TestGenerateCmd_NoFlags(t *testing.T) {
	setup()
	defer teardown()

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.NewGenerateCmd())

	// Set the file argument
	rootCmd.SetArgs([]string{"generate", "testfile.java", "A file to test the test"})

	out := captureOutput(func() {
		err := rootCmd.Execute()
		assert.NoError(t, err)
	})

	assert.Contains(t, out, "No flags provided, defaulting to unit test for file: testfile.java")
	assert.Contains(t, out, "import org.apache.commons.lang3.RandomStringUtils;")
}

// Test case for when only the unit test flag is set
func TestGenerateCmd_UnitTestFlag(t *testing.T) {
	setup()
	defer teardown()

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.NewGenerateCmd())

	// Set the file argument
	rootCmd.SetArgs([]string{"generate", "testfile.java", "--unittest"})

	out := captureOutput(func() {
		err := rootCmd.Execute()
		assert.NoError(t, err)
	})

	assert.Contains(t, out, "Generating unit tests for file: testfile.java")
	assert.Contains(t, out, "import org.apache.commons.lang3.RandomStringUtils;")
}

// Test case for when only the integration test flag is set
func TestGenerateCmd_IntegrationTestFlag(t *testing.T) {
	setup()
	defer teardown()

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.NewGenerateCmd())

	// Set the file argument
	rootCmd.SetArgs([]string{"generate", "testfile.java", "--integrationtest"})

	out := captureOutput(func() {
		err := rootCmd.Execute()
		assert.NoError(t, err)
	})

	assert.Contains(t, out, "Generating integration tests for file: testfile.java")
	assert.Contains(t, out, "import org.apache.commons.lang3.RandomStringUtils;")
}

// Test case for when only the e2e test flag is set
func TestGenerateCmd_E2ETestFlag(t *testing.T) {
	setup()
	defer teardown()

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.NewGenerateCmd())

	// Set the file argument
	rootCmd.SetArgs([]string{"generate", "testfile.java", "--e2etest"})

	out := captureOutput(func() {
		err := rootCmd.Execute()
		assert.NoError(t, err)
	})

	assert.Contains(t, out, "Generating e2e tests for file: testfile.java")
	assert.Contains(t, out, "import org.apache.commons.lang3.RandomStringUtils;")
}

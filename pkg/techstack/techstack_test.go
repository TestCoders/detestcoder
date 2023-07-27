package techstack

import (
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

// Testing SetDependencyManager function to ensure it correctly sets the name and version.
func TestSetDependencyManager(t *testing.T) {
	tStack := NewTechStack()

	tStack.SetDependencyManager("go mod", "1.18")
	assert.Equal(t, "go mod", tStack.TechStack.DependencyManager.Name)
	assert.Equal(t, "1.18", tStack.TechStack.DependencyManager.Version)
}

// Testing SetLanguage function to ensure it correctly sets the name and version.
func TestSetLanguage(t *testing.T) {
	tStack := NewTechStack()

	tStack.SetLanguage("Go", "1.18")
	assert.Equal(t, "Go", tStack.TechStack.Language.Name)
	assert.Equal(t, "1.18", tStack.TechStack.Language.Version)
}

// Testing AddFramework function to ensure it correctly adds a new framework with the name and version.
func TestSetFramework(t *testing.T) {
	tStack := NewTechStack()

	tStack.AddFramework("cobra", "1.0")
	assert.Equal(t, "cobra", tStack.TechStack.Framework.Name)
	assert.Equal(t, "1.0", tStack.TechStack.Framework.Version)
}

// Testing AddTestFramework function to ensure it correctly adds a new unit test framework with the name and version.
func TestSetTestFramework(t *testing.T) {
	tStack := NewTechStack()

	tStack.AddTestFramework("testify", "1.7")
	assert.Equal(t, "testify", tStack.TechStack.TestFramework.Name)
	assert.Equal(t, "1.7", tStack.TechStack.TestFramework.Version)
}

// Testing AddTestDependency function to ensure it correctly adds a new unit test dependency with the name and version.
func TestAddTestDependency(t *testing.T) {
	tStack := NewTechStack()

	tStack.AddTestDependency("testify", "1.7")
	assert.Equal(t, "testify", tStack.TechStack.TestDependencies[0].Name)
	assert.Equal(t, "1.7", tStack.TechStack.TestDependencies[0].Version)
}

// This test verifies that a new TechStack can be created,
// populated with data, and correctly serialized to a YAML file.
func TestWriteCurrentTechStack(t *testing.T) {
	techStack := NewTechStack()
	techStack.SetLanguage("Go", "1.18")
	techStack.SetDependencyManager("dep", "1.0.0")
	techStack.AddFramework("cobra", "1.0.0")
	techStack.AddTestFramework("testify", "1.0.0")
	techStack.AddTestDependency("testify", "1.0.0")

	WriteCurrentTechStack(techStack)

	// Check if .techstack.template file exists
	_, err := os.Stat(".techstack")
	assert.NoError(t, err)

	// Clean up after test
	defer os.Remove(".techstack")
}

// This test verifies that the contents of a populated .techstack.template YAML file
// can be correctly read and deserialized into a TechStack struct.
func TestGetCurrentTechStack(t *testing.T) {
	techStack := NewTechStack()
	techStack.SetLanguage("Go", "1.18")
	techStack.SetDependencyManager("dep", "1.0.0")
	techStack.AddFramework("cobra", "1.0.0")
	techStack.AddTestFramework("testify", "1.0.0")
	techStack.AddTestDependency("testify", "1.0.0")

	WriteCurrentTechStack(techStack)

	gotTechStack := GetCurrentTechStack()

	// Verify that the retrieved tech stack matches the written one
	assert.True(t, reflect.DeepEqual(techStack, gotTechStack))

	// Clean up after test
	defer os.Remove(".techstack")
}

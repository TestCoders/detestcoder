package techstack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Testing SetDependencyManager function to ensure it correctly sets the name and version.
func TestSetDependencyManager(t *testing.T) {
	tStack := NewTechStack()

	tStack.SetDependencyManager("go mod", "1.18")
	assert.Equal(t, "go mod", tStack.DependencyManager.Name)
	assert.Equal(t, "1.18", tStack.DependencyManager.Version)
}

// Testing SetLanguage function to ensure it correctly sets the name and version.
func TestSetLanguage(t *testing.T) {
	tStack := NewTechStack()

	tStack.SetLanguage("Go", "1.18")
	assert.Equal(t, "Go", tStack.Language.Name)
	assert.Equal(t, "1.18", tStack.Language.Version)
}

// Testing SetFramework function to ensure it correctly adds a new framework with the name and version.
func TestSetFramework(t *testing.T) {
	tStack := NewTechStack()

	tStack.SetFramework("cobra", "1.0")
	assert.Equal(t, "cobra", tStack.Framework.Name)
	assert.Equal(t, "1.0", tStack.Framework.Version)
}

// Testing SetTestFramework function to ensure it correctly adds a new unit test framework with the name and version.
func TestSetTestFramework(t *testing.T) {
	tStack := NewTechStack()

	tStack.SetTestFramework("testify", "1.7")
	assert.Equal(t, "testify", tStack.TestFramework.Name)
	assert.Equal(t, "1.7", tStack.TestFramework.Version)
}

// Testing AddTestDependency function to ensure it correctly adds a new unit test dependency with the name and version.
func TestAddTestDependency(t *testing.T) {
	tStack := NewTechStack()

	tStack.AddTestDependency("testify", "1.7")
	assert.Equal(t, "testify", tStack.TestDependencies[0].Name)
	assert.Equal(t, "1.7", tStack.TestDependencies[0].Version)
}

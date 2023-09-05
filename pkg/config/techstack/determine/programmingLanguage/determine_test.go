package programmingLanguage

import (
	"github.com/stretchr/testify/assert"
	dm "github.com/testcoders/detestcoder/pkg/constants/project"
	"log"
	"os"
	"testing"
)

func setup(testfile string) {
	createTempTestfile(testfile)
}

func createTempTestfile(testfile string) {
	content := []byte(`<pom></pom>`)
	err := os.WriteFile(testfile, content, 0644)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
}

func teardown(testfile string) {
	err := os.Remove(testfile)
	if err != nil {
		log.Fatalf("Failed to remove %v: %v", testfile, err)
	}
}

func TestGetProgrammingLanguages_Maven(t *testing.T) {
	setup("pom.pom")
	defer teardown("pom.pom")

	depman := GetDependencyManager()
	pl := GetProgrammingLanguages(depman)
	assert.NotNil(t, pl)
	assert.Equal(t, []string{dm.JAVA, dm.SCALA, dm.KOTLIN}, pl)
}

func TestGetProgrammingLanguages_Pip(t *testing.T) {
	setup("requirements.txt")
	defer teardown("requirements.txt")

	depman := GetDependencyManager()
	pl := GetProgrammingLanguages(depman)
	assert.NotNil(t, pl)
	assert.Equal(t, []string{dm.PYTHON}, pl)
}

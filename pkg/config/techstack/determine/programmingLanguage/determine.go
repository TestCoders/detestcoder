package programmingLanguage

import (
	"github.com/testcoders/detestcoder/pkg/constants/project"
	"os"
)

func GetProgrammingLanguages(manager string) []string {
	if languages, ok := project.DependencyManager[manager]; ok {
		return languages
	}
	return nil
}

func GetDependencyManager() string {
	files, err := os.ReadDir(".")
	if err != nil {
		return "Error reading directory: " + err.Error()
	}

	for _, file := range files {
		if !file.IsDir() {
			for manager, filenames := range project.DependencyManagerFile {
				for _, filename := range filenames {
					if file.Name() == filename {
						return manager
					}
				}
			}
		}
	}

	return ""
}

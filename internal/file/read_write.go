package file

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ProcessFiles(sourceFile, testFile string) error {
	sourceContent, err := readFile(sourceFile)
	if err != nil {
		fmt.Printf("Failed to read source file: %v\n", err)
		return err
	}

	testContent := ""
	if testFile != "" {
		testContent, err = readFile(testFile)
		if err != nil {
			fmt.Printf("Failed to read test file: %v\n", err)
			return err
		}
	}

	result := aiAnalyze(sourceContent, testContent)

	outputFile := ""
	if testFile != "" {
		outputFile = filepath.Dir(testFile) + "/AI-" + filepath.Base(testFile)
	} else {
		outputFile = filepath.Dir(sourceFile) + "/AI-" + filepath.Base(sourceFile)
	}

	err = writeFile(outputFile, result)
	if err != nil {
		fmt.Printf("Failed to write result file: %v\n", err)
		return err
	}
	return nil
}

func readFile(filepath string) (string, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func aiAnalyze(sourceContent, testContent string) string {
	// Perform AI analysis on the source and test content
	if testContent != "" {
		return testContent
	}
	return sourceContent
}

func writeFile(filepath, content string) error {
	err := ioutil.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

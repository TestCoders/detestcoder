package file

import (
	"fmt"
	"io/ioutil"
)

func processFiles(sourceFile, testFile string) error {
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

	err = writeFile("result.txt", result)
	if err != nil {
		fmt.Printf("Failed to write result file: %v\n", err)
		return err
	}

	fmt.Println("Analysis completed successfully. Result written to result.txt.")
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
	// Replace this with your actual implementation

	analysisResult := "This is the analysis result."

	return analysisResult
}

func writeFile(filepath, content string) error {
	err := ioutil.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

package files

import (
	"bufio"
	"fmt"
	"github.com/testcoders/detestcoder/internal/ai"
	"github.com/testcoders/detestcoder/pkg/constants/project"
	"github.com/testcoders/detestcoder/pkg/promptbuilder"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ReadContentsOfFileAndAddCodeSnippet(pb *promptbuilder.PromptBuilder, filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	code := string(bytes)
	pb.AddCodeSnippet(code)
}

func WritePromptToFile(data string) {
	timestamp := time.Now().Unix()

	// create the directory if it doesn't exist
	dir := "generatedPrompts/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	filename := fmt.Sprintf("%sprompt_%d", dir, timestamp)

	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

func WriteOutputToFile(response ai.Response, testFileName string) {
	dir := "generatedOutput/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	testFileName = strings.ReplaceAll(testFileName, "../", "")
	outputFile := strings.Split(testFileName, ".")

	// get the base filename
	baseFilename := filepath.Base(testFileName)
	outputFileName := ""

	if outputFile[1] == project.JAVA {
		outputFileName = strings.TrimSuffix(baseFilename, filepath.Ext(baseFilename)) + "Test." + outputFile[1]
	} else if outputFile[1] == project.KOTLIN {
		outputFileName = strings.TrimSuffix(baseFilename, filepath.Ext(baseFilename)) + "Test." + outputFile[1]
	}

	filename := fmt.Sprintf("%s%s", dir, outputFileName)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	content := response.Content
	content = strings.ReplaceAll(content, "```"+outputFile[1], "")
	content = strings.ReplaceAll(content, "```", "")
	content = strings.TrimSpace(content)
	if _, err := w.WriteString(content + "\n\n"); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	if err = w.Flush(); err != nil {
		log.Fatalf("Failed to flush writer: %v", err)
	}
}

// AppendTimestampToFile appends a timestamp to the filename before its extension.
func AppendTimestampToFile(filePath string) string {
	parts := strings.Split(filePath, ".")
	if len(parts) < 2 {
		// The given filePath does not have an extension.
		return filePath
	}
	base := parts[0]                                 // The portion before the extension.
	ext := parts[1]                                  // The file extension.
	timestamp := time.Now().Format("20060102150405") // Format: YYYYMMDDHHMMSS
	return fmt.Sprintf("%s_%s_generatedTest.%s", base, timestamp, ext)
}

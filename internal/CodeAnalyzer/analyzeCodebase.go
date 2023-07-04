//TODO: Test this code on an actual directory containing a codebase of one of the languages
//TODO: Integrating the code into the 'detestcoder' environment using cmd
//TODO: Adding more languages, for instance C#

package codeanalyzer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Definition of the Language 'class', including multiple variables
type Language struct {
	Name         string
	Keyword      string
	Extension    string
	Subdirectory []string
}

// List of the ten most commonly used programming languages and associated unique keywords
// These keywords are unique to the respective languages within the scope of this list
// It is possible that a code base of the language does not contain the keyword
// But a code base of one of the other included languages should never include the keyword
var languages = []Language{
	{"Java", "synchronized", ".java", []string{"/java", "/spring", "/hibernate", "/maven", "/gradle"}},
	{"Python", "numpy", ".py", []string{"/python", "/django", "/flask", "/pipenv", "/poetry"}},
	{"Ruby", "puts", ".rb", []string{"/ruby", "/rails", "/bundler", "/gemfile"}},
	{"C", "scanf", ".c", []string{"/c"}},
	{"C++", "cout", ".cpp", []string{"/cpp"}},
	{"JavaScript", "console.log", ".js", []string{"/javascript", "/node", "/express", "/react", "/angular", "/vue", "/npm", "/yarn"}},
	{"Go", "fmt.Println", ".go", []string{"/go", "/gin", "/echo", "/glide"}},
	{"Swift", "guard", ".swift", []string{"/swift", "/swiftui", "/cocoapods", "/spm"}},
	{"Rust", "match", ".rs", []string{"/rust", "/cargo"}},
	{"PHP", "$this", ".php", []string{"/php", "/laravel", "/symfony", "/codeigniter", "/composer"}},
}

// TODO: Need to move to Main package
func main() {
	dir := "path/to/directory" // Update with the actual directory path

	lang, err := processDirectory(dir)
	if err != nil {
		fmt.Printf("Failed to process directory: %v\n", err)
	} else {
		fmt.Printf("The language is: %s", lang)
	}
}

// processDirectory recursively traverses the directory and processes each file, returns 'empty' when a directory is empty
// TODO: For now the output is printed to the console, want to implement writing to a file or something similar
// TODO: When one codebase contains multiple languages the codeanalyzer throws an error
// TODO: determine vesions, frameworks, dependencies
func processDirectory(dir string) (lang string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("failed to read directory: %v \n", err)
		return "", err
	}

	if len(files) == 0 {
		fmt.Printf("Directory %s is empty.\n", dir)
		return "", nil
	}

	var filesInDirectory []string
	for _, file := range files {
		filesInDirectory = append(filesInDirectory, file.Name())
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			lang, err := processDirectory(filePath)
			if err != nil {
				fmt.Printf("Failed to process subdirectory: %v\n", err)
				return "", err
			} else {
				fmt.Printf("The language is: %s", lang)
			}
		} else {
			lang, err := detectLanguage(filePath, filesInDirectory)
			if err != nil {
				fmt.Printf("Failed to detect language: %v\n", err)
				return "", err
			} else {
				fmt.Printf("File: %s, Language: %s\n", filePath, lang)
				return lang, nil
			}
		}
	}

	return "", errors.New("Oops, something went wrong during the processing of the directory")
}

// detectLanguage determines the language used in a file based on extension, subdirectory, and keywords
// When output is conflicting or missing, an error is thrown including an informative error message
func detectLanguage(filePath string, filesInDirectory []string) (lang string, err error) {
	extension := filepath.Ext(filePath)
	subdirectory := filepath.Dir(filePath)

	extensionLang, err := detectLanguageByExtension(filesInDirectory, extension)
	if err != nil {
		fmt.Printf("Failed to detect language. %s\n", err)
	}
	subdirLang, err := detectLanguageBySubdirectory(subdirectory)
	if err != nil {
		fmt.Printf("Failed to detect language. %s\n", err)
	}
	keywordLang, err := detectLanguageByKeywords(filePath)
	if err != nil {
		fmt.Printf("Failed to detect language. %s\n", err)
	}

	if extensionLang == "" && subdirLang == "" && keywordLang == "" {
		return "", errors.New("Unknown language")
	}

	if extensionLang != "" && subdirLang != "" && keywordLang != "" {
		if extensionLang == subdirLang && subdirLang == keywordLang {
			return extensionLang, nil
		}
	} else {
		return "", fmt.Errorf("Determined language is ambiguous. Extension: %s, Subdirectory: %s, Keyword: %s", extensionLang, subdirLang, keywordLang)
	}
	return "", errors.New("Oops, something went wrong whilst trying to detemine the language of the codebase")
}

// detectLanguageByExtension determines the language based on file extension
// TODO: not sure if going through all the files in the directory is necessary here, or whether it is redundant compared to the code above
func detectLanguageByExtension(filesInDirectory []string, extension string) (lang string, err error) {
	for _, lang := range languages {
		for _, fileName := range filesInDirectory {
			if hasExtension(fileName, extension) {
				return lang.Name, nil
			}
		}
	}
	return "", fmt.Errorf("No matching language found for extension: %s", extension)
}

// detectLanguageBySubdirectory determines the language based on subdirectory name
func detectLanguageBySubdirectory(subdirectory string) (lang string, err error) {
	for _, lang := range languages {
		if language := hasSubdirectory(subdirectory); language != "" {
			return lang.Name, nil
		}
	}
	return "", fmt.Errorf("No matching language found for subdirectory: %s", subdirectory)
}

// detectLanguageByKeywords determines the language based on keywords in the file content
func detectLanguageByKeywords(filePath string) (lang string, err error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Failed to read the file: %s", filePath)
	}

	code := string(content)

	for _, lang := range languages {
		if hasKeyword(code, lang.Keyword) {
			return lang.Name, nil
		}
	}

	return "", fmt.Errorf("No matching language found in: %s, based on keywords", filePath)
}

// hasExtension checks if the subdirectory or file extension matches
// TODO: Not sure if filepath.Ext is correct or redundant, and if and how to use fileName argument
func hasExtension(fileName, extension string) bool {
	for _, lang := range languages {
		if lang.Extension == extension && filepath.Ext(fileName) == extension {
			return true
		}
	}
	return false
}

// hasSubdirectory checks if the subdirectory contains the language name
// When a subdirectory is found the associated language is returned and no other subdirectories are checked
func hasSubdirectory(subdirectory string) string {
	for _, lang := range languages {
		for _, dir := range lang.Subdirectory {
			if strings.Contains(subdirectory, dir) {
				return lang.Name
			}
		}
	}
	return ""
}

// hasKeywords checks if the code contains all the keywords of a given language
// TODO: This could potentially be a very timeconsuming helperfunction, because the whole codebase is searched for the keyword
//
//	and the added value of the keywords might be minimal. Tradeoff between time and value: To Be Determined.
func hasKeyword(code string, keyword string) bool {
	for _, lang := range languages {
		if lang.Keyword == keyword && strings.Contains(code, keyword) {
			return true
		}
	}
	return false
}

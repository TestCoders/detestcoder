package techstack

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type TechStack struct {
	TechStack struct {
		Language struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		} `yaml:"language"`
		DependencyManager struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		} `yaml:"dependency_manager"`
		Framework struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		} `yaml:"framework"`
		TestFramework struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		} `yaml:"test_framework"`
		TestDependencies []struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		} `yaml:"test_dependencies"`
	} `yaml:"tech_stack"`
}

func GetCurrentTechStack() *TechStack {
	// Open the file
	f, err := os.Open(".techstack")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer f.Close()

	// Read the file's content
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Unmarshal the YAML content into a TechStack struct
	var ts TechStack
	err = yaml.Unmarshal(bytes, &ts)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Return the populated struct
	return &ts
}

func WriteCurrentTechStack(ts *TechStack) {
	// Marshal the ts struct into YAML
	bytes, err := yaml.Marshal(ts)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Write the YAML to the .techstack.template file
	err = ioutil.WriteFile(".techstack", bytes, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func NewTechStack() *TechStack {
	return &TechStack{
		TechStack: struct {
			Language struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			} `yaml:"language"`
			DependencyManager struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			} `yaml:"dependency_manager"`
			Framework struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			} `yaml:"framework"`
			TestFramework struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			} `yaml:"test_framework"`
			TestDependencies []struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			} `yaml:"test_dependencies"`
		}{
			Language: struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			}{},
			DependencyManager: struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			}{},
			Framework: struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			}{},
			TestFramework: struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			}{},
			TestDependencies: []struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			}{},
		},
	}
}

func (t *TechStack) SetDependencyManager(name string, version string) {
	t.TechStack.DependencyManager.Name = name
	t.TechStack.DependencyManager.Version = version
}

func (t *TechStack) SetLanguage(name string, version string) {
	t.TechStack.Language.Name = name
	t.TechStack.Language.Version = version
}

func (t *TechStack) AddFramework(name string, version string) {
	t.TechStack.Framework.Name = name
	t.TechStack.Framework.Version = version
}

func (t *TechStack) AddTestFramework(name string, version string) {
	t.TechStack.TestFramework.Name = name
	t.TechStack.TestFramework.Version = version
}

func (t *TechStack) AddTestDependency(name string, version string) {
	unitTestDependency := struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}{
		Name:    name,
		Version: version,
	}
	t.TechStack.TestDependencies = append(t.TechStack.TestDependencies, unitTestDependency)
}

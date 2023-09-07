package project

const (
	MAVEN  = "maven"
	GRADLE = "gradle"
	SDT    = "sdt"
	PIP    = "pip"
	CONDA  = "conda"
	NPM    = "npm"
	YARN   = "yarn"
)

var DependencyManager = map[string][]string{
	MAVEN:  {JAVA, SCALA, KOTLIN},
	GRADLE: {JAVA, SCALA, KOTLIN},
	SDT:    {SCALA},
	PIP:    {PYTHON},
	CONDA:  {PYTHON},
	NPM:    {JAVASCRIPT, TYPESCRIPT},
	YARN:   {JAVASCRIPT, TYPESCRIPT},
}

var DependencyManagerFile = map[string][]string{
	MAVEN:  {"pom.xml"},
	GRADLE: {"build.gradle", "build.gradle.kts"},
	SDT:    {"build.sbt"},
	PIP:    {"requirements.txt"},
	CONDA:  {"environment.yml"},
	NPM:    {"package.json"},
	YARN:   {"package.json", "yarn.lock"},
}

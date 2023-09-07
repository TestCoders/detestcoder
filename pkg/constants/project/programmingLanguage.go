package project

const (
	JAVA       = "java"
	SCALA      = "scala"
	KOTLIN     = "kotlin"
	PYTHON     = "python"
	JAVASCRIPT = "javascript"
	TYPESCRIPT = "typescript"
)

var LanguageExtensions = map[string][]string{
	JAVA:       {"java"},
	SCALA:      {"scala"},
	KOTLIN:     {"kt"},
	PYTHON:     {"py"},
	JAVASCRIPT: {"js", "jsx"},
	TYPESCRIPT: {"ts", "tsx"},
}

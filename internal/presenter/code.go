package presenter

import (
	"bytes"
	"github.com/z6wdc/go-leetcode/internal/entity"
	"strings"
	"text/template"
	"unicode"
)

// Template for <slug>.go
const solutionTemplate = `package problems

func {{ .FuncName }}() {
	// TODO: implement
}
`

// Template for <slug>_test.go
const testTemplate = `package problems

import "testing"

func Test{{ .FuncName }}(t *testing.T) {
	// TODO: write test
}
`

type CodeTemplateData struct {
	FuncName string
}

func RenderSolutionCode(q *entity.Question) (filename string, content string, err error) {
	funcName := slugToCamel(q.Slug)
	snakeName := slugToSnake(q.Slug) + ".go"

	tmpl := template.Must(template.New("solution").Parse(solutionTemplate))
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, CodeTemplateData{FuncName: funcName})
	if err != nil {
		return "", "", err
	}

	return snakeName, buf.String(), nil
}

func RenderTestCode(q *entity.Question) (filename string, content string, err error) {
	funcName := slugToCamel(q.Slug)
	snakeName := slugToSnake(q.Slug) + "_test.go"

	tmpl := template.Must(template.New("test").Parse(testTemplate))
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, CodeTemplateData{FuncName: funcName})
	if err != nil {
		return "", "", err
	}

	return snakeName, buf.String(), nil
}

func slugToCamel(slug string) string {
	parts := strings.Split(slug, "-")
	for i, part := range parts {
		parts[i] = capitalize(part)
	}
	return strings.Join(parts, "")
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func slugToSnake(slug string) string {
	return strings.ReplaceAll(slug, "-", "_")
}

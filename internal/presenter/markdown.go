package presenter

import (
	"bytes"
	"github.com/z6wdc/go-leetcode/internal/entity"
	"text/template"
)

const readmeTemplate = `# {{ .Title }}

**Difficulty**: {{ .Difficulty }}

**Tags**: {{ range .Tags }}` + "`{{ . }}` " + `{{ end }}

---

{{ .Content }}
`

func RenderMarkdown(q *entity.Question) (string, error) {
	tmpl, err := template.New("readme").Parse(readmeTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, q); err != nil {
		return "", err
	}

	return buf.String(), nil
}

package presenter

import (
	"bytes"
	"github.com/z6wdc/go-leetcode/internal/entity"
	"github.com/z6wdc/go-leetcode/internal/util"
	"text/template"
)

const readmeTemplate = `# {{ .Title }}

**Difficulty**: {{ .Difficulty }}

**Tags**: {{ range .Tags }}` + "`{{ . }}` " + `{{ end }}

---

{{ .Content }}
`

type markdownData struct {
	Title      string
	Difficulty string
	Tags       []string
	Content    string
}

func RenderMarkdown(q *entity.Question) (string, error) {
	tmpl, err := template.New("readme").Parse(readmeTemplate)
	if err != nil {
		return "", err
	}

	data := markdownData{
		Title:      q.Title,
		Difficulty: q.Difficulty,
		Tags:       q.Tags,
		Content:    util.CleanText(q.Content),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

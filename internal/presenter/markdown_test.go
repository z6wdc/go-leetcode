package presenter_test

import (
	"github.com/z6wdc/go-leetcode/internal/entity"
	"github.com/z6wdc/go-leetcode/internal/presenter"
	"strings"
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	q := &entity.Question{
		ID:         1,
		Title:      "Two Sum",
		Slug:       "two-sum",
		Content:    "Given\u00A0an array of integers...",
		Difficulty: "Easy",
		Tags:       []string{"Array", "Hash Table"},
	}

	md, err := presenter.RenderMarkdown(q)
	if err != nil {
		t.Fatalf("RenderMarkdown failed: %v", err)
	}

	if !strings.Contains(md, "# Two Sum") {
		t.Errorf("expected title header in markdown")
	}
	if !strings.Contains(md, "**Difficulty**: Easy") {
		t.Errorf("expected difficulty line")
	}
	if !strings.Contains(md, "`Array`") || !strings.Contains(md, "`Hash Table`") {
		t.Errorf("expected tags")
	}
	if strings.Contains(md, "\u00A0") {
		t.Error("Markdown still contains NBSP")
	}
	t.Logf("\nGenerated Markdown:\n%s", md)
}

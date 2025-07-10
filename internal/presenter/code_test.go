package presenter_test

import (
	"github.com/z6wdc/go-leetcode/internal/entity"
	"github.com/z6wdc/go-leetcode/internal/presenter"
	"strings"
	"testing"
)

func TestRenderSolutionCode(t *testing.T) {
	q := &entity.Question{
		Title: "Two Sum",
		Slug:  "two-sum",
	}

	filename, content, err := presenter.RenderSolutionCode(q)
	if err != nil {
		t.Fatalf("RenderSolutionCode failed: %v", err)
	}

	if filename != "two_sum.go" {
		t.Errorf("expected filename 'two_sum.go', got '%s'", filename)
	}
	if !strings.Contains(content, "func TwoSum") {
		t.Errorf("expected content to include 'func TwoSum'")
	}
	if !strings.Contains(content, "package problems") {
		t.Errorf("expected content to include 'package problems'")
	}
}

func TestRenderTestCode(t *testing.T) {
	q := &entity.Question{
		Title: "Two Sum",
		Slug:  "two-sum",
	}

	filename, content, err := presenter.RenderTestCode(q)
	if err != nil {
		t.Fatalf("RenderTestCode failed: %v", err)
	}

	if filename != "two_sum_test.go" {
		t.Errorf("expected filename 'two_sum_test.go', got '%s'", filename)
	}
	if !strings.Contains(content, "TestTwoSum") {
		t.Errorf("expected content to include 'TestTwoSum'")
	}
	if !strings.Contains(content, "package problems") {
		t.Errorf("expected content to include 'package problems'")
	}
}

package leetcode_test

import (
	"github.com/z6wdc/go-leetcode/internal/adapter/leetcode"
	"testing"
)

func TestFetchQuestionBySlug(t *testing.T) {
	slug := "two-sum"

	q, err := leetcode.FetchQuestionBySlug(slug)
	if err != nil {
		t.Fatalf("failed to fetch question: %v", err)
	}

	if q.ID != 1 {
		t.Errorf("expected ID 1, got %d", q.ID)
	}
	if q.Title != "Two Sum" {
		t.Errorf("expected title 'Two Sum', got '%s'", q.Title)
	}
	if q.Slug != "two-sum" {
		t.Errorf("expected slug 'two-sum', got '%s'", q.Slug)
	}
	if q.Difficulty != "Easy" {
		t.Errorf("expected difficulty 'Easy', got '%s'", q.Difficulty)
	}
	if q.Content == "" {
		t.Errorf("expected non-empty content")
	}
	if len(q.Tags) == 0 {
		t.Errorf("expected at least one tag")
	}
}

package usecase

import (
	"fmt"
	"github.com/z6wdc/go-leetcode/internal/adapter/leetcode"
	"github.com/z6wdc/go-leetcode/internal/infra/writer"
	"github.com/z6wdc/go-leetcode/internal/presenter"
	"path/filepath"
)

func GenerateQuestionBySlug(slug string) error {
	q, err := leetcode.FetchQuestionBySlug(slug)
	if err != nil {
		return fmt.Errorf("fetch question failed: %w", err)
	}

	readme, err := presenter.RenderMarkdown(q)
	if err != nil {
		return fmt.Errorf("render markdown failed: %w", err)
	}

	solutionName, solutionCode, err := presenter.RenderSolutionCode(q)
	if err != nil {
		return fmt.Errorf("render solution code failed: %w", err)
	}

	testName, testCode, err := presenter.RenderTestCode(q)
	if err != nil {
		return fmt.Errorf("render test code failed: %w", err)
	}

	baseDir := filepath.Join("generated", "problems", q.Slug)

	if err := writer.WriteFile(baseDir, "README.md", readme); err != nil {
		return err
	}
	if err := writer.WriteFile(baseDir, solutionName, solutionCode); err != nil {
		return err
	}
	if err := writer.WriteFile(baseDir, testName, testCode); err != nil {
		return err
	}

	return nil
}

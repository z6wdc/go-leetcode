package writer_test

import (
	"github.com/z6wdc/go-leetcode/internal/infra/writer"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFile(t *testing.T) {
	baseDir := "test_output"
	dir := filepath.Join(baseDir, "example-problem")
	filename := "sample.txt"
	content := "Hello, Writer!"

	err := writer.WriteFile(dir, filename, content)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	fullPath := filepath.Join(dir, filename)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("failed to read file back: %v", err)
	}

	if string(data) != content {
		t.Errorf("expected content '%s', got '%s'", content, string(data))
	}

	err = os.RemoveAll(baseDir)
	if err != nil {
		return
	}
}

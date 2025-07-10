package util_test

import (
	"github.com/z6wdc/go-leetcode/internal/util"
	"testing"
)

func TestCleanText(t *testing.T) {
	input := "Hello\u00A0World"
	expected := "Hello World"
	output := util.CleanText(input)

	if output != expected {
		t.Errorf("expected '%s', got '%s'", expected, output)
	}
}

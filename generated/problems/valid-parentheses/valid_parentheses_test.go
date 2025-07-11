package problems

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"{[()]}", true},
		{"(([]){})", true},

		{"(]", false},
		{"([)]", false},
		{"(", false},
		{"]", false},
		{"([]", false},
		{"{[(])}", false},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result := isValid(tc.input)
			if result != tc.expected {
				t.Errorf("isValid(%q) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

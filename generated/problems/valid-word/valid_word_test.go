package problems

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		word     string
		expected bool
	}{
		// Valid: length >= 3, contains at least one vowel and one consonant, only letters or digits
		{"abc", true},
		{"a1c", true},
		{"A2C", true},
		{"ab2", true},

		// Invalid: length less than 3
		{"a", false},
		{"ab", false},

		// Invalid: contains illegal characters
		{"a!c", false},
		{"a_c", false},
		{"a.c", false},

		// Invalid: only vowels
		{"aei", false},
		{"AEI", false},

		// Invalid: only consonants
		{"bcd", false},
		{"BCD", false},

		// Invalid: only digits
		{"123", false},

		// Invalid: digits + vowels, but no consonants
		{"a1e", false},

		// Invalid: digits + consonants, but no vowels
		{"1bc", false},
	}

	for _, test := range tests {
		t.Run(test.word, func(t *testing.T) {
			result := isValid(test.word)
			if result != test.expected {
				t.Errorf("isValid(%q) = %v; expected %v", test.word, result, test.expected)
			}
		})
	}
}

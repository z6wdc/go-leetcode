package problems

import "unicode"

func isValid(word string) bool {
	// Check if the word is empty
	if len(word) < 3 {
		return false
	}

	hasVowel := false
	hasConsonant := false

	for _, r := range word {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false // Contains illegal characters
		}

		if unicode.IsLetter(r) {
			// Check if the character is a vowel or consonant
			u := unicode.ToLower(r)
			if u == 'a' || u == 'e' || u == 'i' || u == 'o' || u == 'u' {
				hasVowel = true // Found a vowel
			} else {
				hasConsonant = true // Found a consonant
			}
		}
	}
	return hasVowel && hasConsonant
}

package util

import "strings"

func CleanText(s string) string {
	s = strings.ReplaceAll(s, "\u00A0", " ")

	return s
}

package utils

import (
	"strings"
	"unicode"
)

func trimString(input string) string {
	// Trim leading and trailing spaces
	trimmed := strings.TrimSpace(input)

	// Replace consecutive spaces with a single space
	var result strings.Builder
	prevChar := ' '
	for _, char := range trimmed {
		if !(unicode.IsSpace(prevChar) && unicode.IsSpace(char)) {
			result.WriteRune(char)
		}
		prevChar = char
	}

	return result.String()
}

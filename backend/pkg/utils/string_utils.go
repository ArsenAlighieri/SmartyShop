package utils

import (
	"strings"
)

// ConvertToEnglishChars converts Turkish characters in a string to their English equivalents.
func ConvertToEnglishChars(s string) string {
	replacer := strings.NewReplacer(
		"ç", "c", "Ç", "C",
		"ğ", "g", "Ğ", "G",
		"ı", "i", "İ", "I",
		"ö", "o", "Ö", "O",
		"ş", "s", "Ş", "S",
		"ü", "u", "Ü", "U",
	)
	return replacer.Replace(s)
}

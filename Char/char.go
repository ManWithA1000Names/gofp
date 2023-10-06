package Char

import "unicode"

// Detect upper case unicode characters.
func IsUpper(c rune) bool {
	return unicode.IsUpper(c)
}

// Detect lower case unicode characters.
func IsLower(c rune) bool {
	return unicode.IsLower(c)
}

// Detect upper case and lower case unicode characters.
func IsAlpha(c rune) bool {
	return unicode.IsLetter(c)
}

// Detect digits 0123456789
func IsDigit(c rune) bool {
	return unicode.IsDigit(c)
}

// Detect upper case and lower case unicode characters.
func IsAlphaNum(c rune) bool {
	return IsAlpha(c) || IsDigit(c)
}

// Detect space unicode characters.
func IsSpace(c rune) bool {
	return unicode.IsSpace(c)
}

// Convert to upper case.
func ToUpper(c rune) rune {
	return unicode.ToUpper(c)
}

// Convert to lower case.
func ToLower(c rune) rune {
	return unicode.ToLower(c)
}

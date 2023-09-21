package Char

import "unicode"

func IsUpper(c rune) bool {
	return unicode.IsUpper(c)
}

func IsLower(c rune) bool {
	return unicode.IsLower(c)
}

func IsAlpha(c rune) bool {
	return unicode.IsLetter(c)
}

func IsDigit(c rune) bool {
	return unicode.IsDigit(c)
}

func IsAlphaNum(c rune) bool {
	return IsAlpha(c) || IsDigit(c)
}

func IsSpace(c rune) bool {
	return unicode.IsSpace(c)
}

func ToUpper(c rune) rune {
	return unicode.ToUpper(c)
}

func ToLower(c rune) rune {
	return unicode.ToLower(c)
}

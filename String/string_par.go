package String

import "github.com/manwitha1000names/gofp/v2/List"

// HIGHER-ORDER FUNCTIONS

// Transform every character in a string
func Map_par(mapfn func(c rune) rune, s string) string {
	return FromList(List.Map_par(mapfn, ToList(s)))
}

// Keep only the characters that pass the test.
// This MESSES WITH THE ORDERING of the characters!
// Only use this when you are using a string purely as a set of characters.
func Filter_par(testfn func(c rune) bool, s string) string {
	return FromList(List.Filter_par(testfn, ToList(s)))
}

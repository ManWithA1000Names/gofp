package String

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/manwitha1000names/gofp/v3/Basics"
	"github.com/manwitha1000names/gofp/v3/Char"
	"github.com/manwitha1000names/gofp/v3/List"
	. "github.com/manwitha1000names/gofp/v3/MaybeResult"
	"github.com/manwitha1000names/gofp/v3/Tuple"
)

// Determine if a string is empty.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// Reverse a string.
func Reverse(s string) string {
	return FromList(List.Reverse(ToList(s)))
}

// Repeat a string n times.
func Repeat(amount int, s string) string {
	return Concat(List.Repeat(amount, s))
}

// Replace all occurrences of some substring.
func Replace(this string, that string, s string) string {
	return strings.ReplaceAll(s, this, that)
}

// BUILDING AND SPLITTING

// Append two strings.
func Append(s string, s1 string) string {
	return s + s1
}

// Concatenate many strings into one.
func Concat(list []string) string {
	return List.Foldr(Append, "", list)
}

// Split a string using a given separator.
func Split(delimiter string, s string) []string {
	return strings.Split(s, delimiter)
}

// Put many strings together with a given separator.
func Join(delimiter string, list []string) string {
	return strings.Join(list, delimiter)
}

// Break a string into words, splitting on chunks of whitespace.
func Words(s string) []string {
	length := 0
	words := make([]string, 0, 8)
	for i, c := range s {
		if !Char.IsSpace(c) {
			length += 1
		} else if length > 0 {
			words = append(words, s[i-length:i])
			length = 0
		}
	}
	return words
}

// Break a string into lines, splitting on newlines.
func Lines(s string) []string {
	return strings.Split(s, "\n")
}

// GET SUBSTRINGS

// Take a substring given a start and end index.
// Negative indexes are taken starting from the end of the list.
func Slice(start int, stop int, s string) string {
	if start < 0 {
		start = Basics.Max(len(s)+start, 0)
	}
	if stop < 0 {
		stop = Basics.Min(len(s)+stop, len(s))
	}
	if stop <= 0 || start >= stop {
		return ""
	}
	return s[start:stop]
}

// Take n characters from the left side of a string.
func Left(n int, s string) string {
	if n < 0 {
		return ""
	}
	if n > len(s) {
		return s
	}
	return s[0:n]
}

// Take n characters from the right side of a string.
func Right(n int, s string) string {
	if n < 0 {
		return ""
	}
	length := len(s)
	if n >= length {
		return s
	}
	return s[length-n:]
}

// Drop n characters from the left side of a string.
func DropLeft(n int, s string) string {
	if n < 0 {
		return s
	}
	length := len(s)
	if n >= length {
		return ""
	}
	return s[n:]
}

// Drop n characters from the right side of a string.
func DropRight(n int, s string) string {
	if n < 0 {
		return s
	}
	length := len(s)
	if n >= length {
		return ""
	}
	return s[0:n]
}

// See if the second string contains the first one.
func Contains(sub string, s string) bool {
	return strings.Contains(s, sub)
}

// See if the second string starts with the first one.
func StartsWith(sub string, s string) bool {
	length_s := len(s)
	length_sub := len(sub)
	if length_sub > length_s {
		return false
	}
	return sub == s[0:length_sub]
}

// See if the second string ends with the first one.
func EndsWith(sub string, s string) bool {
	length_s := len(s)
	length_sub := len(sub)
	if length_sub > length_s {
		return false
	}
	return sub == s[length_s-length_sub:]
}

// Get all of the indexes for a substring in another string.
func Indeces(sub string, s string) []int {
	length_sub := len(sub)
	indeces := make([]int, 0)
	for i := 0; i <= (len(s) - length_sub); i++ {
		if sub == s[i:i+length_sub] {
			indeces = append(indeces, i)
		}
	}
	return indeces
}

// INT CONVERSIONS

// Try to convert a string into an int, failing on improperly formatted strings.
func ToInt(s string) Maybe[int] {
	i, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return Nothing[int]()
	}
	return Just(int(i))
}

// Convert an Int to a String.
func FromInt[N Basics.Int](i N) string {
	return fmt.Sprint(i)
}

// FLOAT CONVERSIONS

// Try to convert a string into a float, failing on improperly formatted strings.
func ToFloat(s string) Maybe[float64] {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return Nothing[float64]()
	}
	return Just(i)
}

// Convert a Float to a String.
func FromFloat[F Basics.Float](f F) string {
	return fmt.Sprint(f)
}

// CHAR CONVERSIONS

// Create a string from a given character.
func FromChar(c rune) string {
	return string([]rune{c})
}

// Add a character to the beginning of a string.
func Cons(c rune, s string) string {
	return FromChar(c) + s
}

// Split a non-empty string into its head and tail.
// This lets you pattern match on strings exactly as you would with lists.
func Uncons(s string) Maybe[Tuple.Tuple[rune, string]] {
	if IsEmpty(s) {
		return Nothing[Tuple.Tuple[rune, string]]()
	}
	return Just(Tuple.Pair(ToList(s)[0], s[1:]))
}

// LIST CONVERSIONS

// Convert a string to a list of characters.
func ToList(s string) []rune {
	return []rune(s)
}

// Convert a list of characters into a String.
// Can be useful if you want to create a string primarily by consing, perhaps for decoding something.
func FromList(list []rune) string {
	return string(list)
}

// FORMATTING

// Convert a string to all upper case. Useful for case-insensitive comparisons and VIRTUAL YELLING.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Convert a string to all lower case. Useful for case-insensitive comparisons.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Convert the first character of a string to capitali case if it isn't already.
func Capitalize(s string) string {
	runes := ToList(s)
	if len(runes) == 0 || unicode.IsUpper(runes[0]) {
		return s
	}
	runes[0] = unicode.ToUpper(runes[0])
	return FromList(runes)
}

// Pad a string on both sides until it has a given length.
func Pad(new_len int, c rune, s string) string {
	length := len(s)
	half := Basics.ToFloat(new_len-length) / 2
	char_s := FromChar(c)
	return Repeat(Basics.Ceiling[int](half), char_s) + s + Repeat(Basics.Floor[int](half), char_s)
}

// Pad a string on the left until it has a given length.
func PadLeft(new_len int, c rune, s string) string {
	return Repeat(new_len-len(s), FromChar(c)) + s
}

// Pad a string on the right until it has a given length.
func PadRight(new_len int, c rune, s string) string {
	return s + Repeat(new_len-len(s), FromChar(c))
}

// Get rid of whitespace on both sides of a string.
func Trim(s string) string {
	return strings.TrimFunc(s, Char.IsSpace)
}

// Get rid of whitespace on the left of a string.
func TrimLeft(s string) string {
	return strings.TrimLeftFunc(s, Char.IsSpace)
}

// Get rid of whitespace on the right of a string.
func TrimRight(s string) string {
	return strings.TrimRightFunc(s, Char.IsSpace)
}

// HIGHER-ORDER FUNCTIONS

// Transform every character in a string
func Map(mapfn func(c rune) rune, s string) string {
	return FromList(List.Map(mapfn, ToList(s)))
}

// Keep only the characters that pass the test.
func Filter(testfn func(c rune) bool, s string) string {
	return FromList(List.Filter(testfn, ToList(s)))
}

// Reduce a string from the left.
func Foldl[Acc any](reducer func(c rune, acc Acc) Acc, init Acc, s string) Acc {
	return List.Foldl(reducer, init, ToList(s))
}

// Reduce a string from the right.
func Foldr[Acc any](reducer func(c rune, acc Acc) Acc, init Acc, s string) Acc {
	return List.Foldr(reducer, init, ToList(s))
}

// Determine whether any characters pass the test.
func Any(testfn func(c rune) bool, s string) bool {
	return List.Any(testfn, ToList(s))
}

// Determine whether all characters pass the test.
func All(testfn func(c rune) bool, s string) bool {
	return List.All(testfn, ToList(s))
}

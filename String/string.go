package String

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/manwitha1000names/gofp/Basics"
	"github.com/manwitha1000names/gofp/Char"
	"github.com/manwitha1000names/gofp/List"
	"github.com/manwitha1000names/gofp/Maybe"
	"github.com/manwitha1000names/gofp/Tuple"
)

func IsEmpty(s string) bool {
	return len(s) == 0
}

func Reverse(s string) string {
	return FromList(List.Reverse(ToList(s)))
}

func Repeat(amount int, s string) string {
	return Concat(List.Repeat(amount, s))
}

func Replace(this string, that string, s string) string {
	return strings.ReplaceAll(s, this, that)
}

// BUILDING AND SPLITTING

func Append(s string, s1 string) string {
	return s + s1
}

func Concat(list []string) string {
	return List.Foldr(Append, "", list)
}

func Split(delimiter string, s string) []string {
	return strings.Split(s, delimiter)
}

func Join(delimiter string, list []string) string {
	return strings.Join(list, delimiter)
}

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

func Lines(s string) []string {
	return strings.Split(s, "\n")
}

// GET SUBSTRINGS

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

func Left(n int, s string) string {
	if n < 0 {
		return ""
	}
	if n > len(s) {
		return s
	}
	return s[0:n]
}

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

func Contains(sub string, s string) bool {
	return strings.Contains(s, sub)
}

func StartsWith(sub string, s string) bool {
	length_s := len(s)
	length_sub := len(sub)
	if length_sub > length_s {
		return false
	}
	return sub == s[0:length_sub]
}

func EndsWith(sub string, s string) bool {
	length_s := len(s)
	length_sub := len(sub)
	if length_sub > length_s {
		return false
	}
	return sub == s[length_s-length_sub:]
}

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

func ToInt(s string) Maybe.Maybe[int] {
	i, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return Maybe.Nothing[int]()
	}
	return Maybe.Just(int(i))
}

func FromInt[N Basics.Int](i N) string {
	return fmt.Sprint(i)
}

// FLOAT CONVERSIONS

func ToFloat(s string) Maybe.Maybe[float64] {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return Maybe.Nothing[float64]()
	}
	return Maybe.Just(i)
}

func FromFloat[F Basics.Float](f F) string {
	return fmt.Sprint(f)
}

// CHAR CONVERSIONS

func FromChar(c rune) string {
	return string([]rune{c})
}

func Cons(c rune, s string) string {
	return FromChar(c) + s
}

func Uncons(s string) Maybe.Maybe[Tuple.Tuple[rune, string]] {
	if IsEmpty(s) {
		return Maybe.Nothing[Tuple.Tuple[rune, string]]()
	}
	return Maybe.Just(Tuple.Pair(ToList(s)[0], s[1:]))
}

// LIST CONVERSIONS

func ToList(s string) []rune {
	return []rune(s)
}

func FromList(list []rune) string {
	return string(list)
}

// FORMATTING

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func Capitalize(s string) string {
	runes := ToList(s)
  if len(runes) == 0 || unicode.IsUpper(runes[0]) {
    return s
  }
	runes[0] = unicode.ToUpper(runes[0])
	return FromList(runes)
}

func Pad(new_len int, c rune, s string) string {
	length := len(s)
	half := Basics.ToFloat(new_len-length) / 2
	char_s := FromChar(c)
	return Repeat(Basics.Ceiling[int](half), char_s) + s + Repeat(Basics.Floor[int](half), char_s)
}

func PadLeft(new_len int, c rune, s string) string {
	return Repeat(new_len-len(s), FromChar(c)) + s
}

func PadRight(new_len int, c rune, s string) string {
	return s + Repeat(new_len-len(s), FromChar(c))
}

func Trim(s string) string {
	return strings.TrimFunc(s, Char.IsSpace)
}

func TrimLeft(s string) string {
	return strings.TrimLeftFunc(s, Char.IsSpace)
}

func TrimRight(s string) string {
	return strings.TrimRightFunc(s, Char.IsSpace)
}

// HIGHER-ORDER FUNCTIONS

func Map(mapfn func(c rune) rune, s string) string {
	return FromList(List.Map(mapfn, ToList(s)))
}

func Filter(testfn func(c rune) bool, s string) string {
	return FromList(List.Filter(testfn, ToList(s)))
}

func Foldl[Acc any](reducer func(c rune, acc Acc) Acc, init Acc, s string) Acc {
	return List.Foldl(reducer, init, ToList(s))
}

func Foldr[Acc any](reducer func(c rune, acc Acc) Acc, init Acc, s string) Acc {
	return List.Foldr(reducer, init, ToList(s))
}

func Any(testfn func(c rune) bool, s string) bool {
	return List.Any(testfn, ToList(s))
}

func All(testfn func(c rune) bool, s string) bool {
	return List.All(testfn, ToList(s))
}

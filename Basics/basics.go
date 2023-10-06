package Basics

import (
	"math"
)

const (
	LT = -1
	EQ = 0
	GT = 1

	E  = math.E
	PI = math.Pi
)

type Order = int

type SignedInt interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type UnsingnedInt interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr
}

type Int interface {
	SignedInt | UnsingnedInt
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Int | Float
}

type SignedNumbers interface {
	SignedInt | Float
}

type Ordered interface {
	Int | Float | ~string
}

// MATH

// Add two numbers.
func Add[T Number](a, b T) T {
	return a + b
}

// Subtract two numbers.
func Sub[T Number](a, b T) T {
	return a - b
}

// Multiply two numbers
func Mul[T Number](a, b T) T {
	return a * b
}

// Floating-point division
func Div[T Number](a, b T) T {
	return a / b
}

// Integer division
func IDiv[T Int](a, b T) T {
	return a / b
}

// Exponentiation
func Pow[T Number](base, exp T) float64 {
	return math.Pow(float64(base), float64(exp))
}

// Int to Float / Float to Int

// Convert an integer into a float.
func ToFloat[T Int](num T) float64 {
	return float64(num)
}

// Round a number to the nearest integer.
func Round[I Int, T Float](num T) I {
	return I(math.Round(float64(num)))
}

// Floor function, rounding down.
func Floor[I Int, T Float](num T) I {
	return I(math.Floor(float64(num)))
}

// Ceiling function, rounding up.
func Ceiling[I Int, T Float](num T) I {
	return I(math.Ceil(float64(num)))
}

// Truncate a number, rounding towards zero.
func Truncate[I Int, T Float](num T) I {
	return I(num)
}

// EQUALITY

// Check if values are “the same”.
func Eq[T comparable](v1, v2 T) bool {
	return v1 == v2
}

// Check if values are not “the same”.
func Ne[T comparable](v1, v2 T) bool {
	return v1 != v2
}

// COMPARISON

// Check if a value is "less than" (lt) than the other
func Lt[T Ordered](v1, v2 T) bool {
	return v1 < v2
}

// Check if a value is "less than or equal" (le) than the other
func Le[T Ordered](v1, v2 T) bool {
	return v1 <= v2
}

// Check if a value is "greater than" (gt) than the other
func Gt[T Ordered](v1, v2 T) bool {
	return v1 > v2
}

// Check if a value is "greater than or equal" (ge) than the other
func Ge[T Ordered](v1, v2 T) bool {
	return v1 >= v2
}

// Find the larger of all the ordered values.
func Max[T Ordered](v1 T, v ...T) T {
	m := v1
	for _, val := range v {
		if val > m {
			m = val
		}
	}
	return m
}

// Find the smaller of all the ordered values.
func Min[T Ordered](v1 T, v ...T) T {
	m := v1
	for _, val := range v {
		if val < m {
			m = val
		}
	}
	return m
}

// Compare any two comparable values.
func Compare[T Ordered](v1 T, v2 T) Order {
	if v1 == v2 {
		return EQ
	} else if v1 < v2 {
		return LT
	} else {
		return GT
	}
}

// BOOLEANS

// Negate a boolean value.
func Not(v bool) bool {
	return !v
}

// The logical AND operator. True if both inputs are True.
func And(v1, v2 bool) bool {
	return v1 && v2
}

// The logical OR operator. True if one or both inputs are True.
func Or(v1, v2 bool) bool {
	return v1 || v2
}

// The exclusive-or operator. True if exactly one input is True.
func Xor(v1, v2 bool) bool {
	return v1 != v2
}

// APPEND STRINGS

// Put two appendable strings together.
func AppendStrings(v1, v2 string) string {
	return v1 + v2
}

// FANCIER MATH

// Perform modular arithmetic.
// Our modBy function works in the typical mathematical way when you run into negative numbers.
// Use remainderBy for a different treatment of negative numbers,
// or read Daan Leijen’s Division and Modulus for Computer Scientists for more information.
func ModBy[T Int](by, base T) T {
	answer := base % by
	if (answer > 0 && by < 0) || answer < 0 && by > 0 {
		return answer + by
	}
	return answer
}

// Get the remainder after division.
// Use modBy for a different treatment of negative numbers,
// or read Daan Leijen’s Division and Modulus for Computer Scientists for more information.
func RemainderBy[T Int](by, base T) T {
	return base % by
}

// Negate a number.
func Negate[T SignedNumbers](n T) T {
	return n * (T)(-1)
}

// Clamps a number within a given range.
func Clamp[T Number](low, n, high T) T {
	if n < low {
		return low
	}
	if n > high {
		return high
	}
	return n
}

// Calculate the logarithm of a number with a given base.
func LogBase[T Number](base, value T) T {
	return T(math.Log(float64(value)) / math.Log(float64(base)))
}

// FUNCTION HELPERS

// Given a value, returns exactly the same value. This is called the identity function.
func Identity[T any](value T) T { return value }

// Create a function that always returns the same value.
func Always[T any, U any](a T, b U) T { return a }

// Apply the function with the argument given, taking the function as the first argument.
func ApplyL[T, U any](f func(value T) U, val T) U {
	return f(val)
}

// Apply the function with the argument given, taking the function as the second argument.
func ApplyR[T, U any](val T, f func(value T) U) U {
	return f(val)
}

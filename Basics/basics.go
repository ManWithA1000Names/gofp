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

func Add[T Number](a, b T) T {
	return a + b
}

func Sub[T Number](a, b T) T {
	return a - b
}

func Mul[T Number](a, b T) T {
	return a * b
}

func Div[T Number](a, b T) T {
	return a / b
}

func IDiv[T Int](a, b T) T {
	return a / b
}

func Pow[T Number](base, exp T) float64 {
	return math.Pow(float64(base), float64(exp))
}

func ToFloat[T Int](num T) float64 {
	return float64(num)
}

func Round[I Int, T Float](num T) I {
	return I(math.Round(float64(num)))
}

func Floor[I Int, T Float](num T) I {
	return I(math.Floor(float64(num)))
}

func Ceiling[I Int, T Float](num T) I {
	return I(math.Ceil(float64(num)))
}

func Truncate[I Int, T Float](num T) I {
	return I(num)
}

func Eq[T comparable](v1, v2 T) bool {
	return v1 == v2
}

func Ne[T comparable](v1, v2 T) bool {
	return v1 != v2
}

func Lt[T Ordered](v1, v2 T) bool {
	return v1 < v2
}

func Le[T Ordered](v1, v2 T) bool {
	return v1 <= v2
}

func Gt[T Ordered](v1, v2 T) bool {
	return v1 > v2
}

func Ge[T Ordered](v1, v2 T) bool {
	return v1 >= v2
}

func Max[T Ordered](v1 T, v ...T) T {
	m := v1
	for _, val := range v {
		if val > m {
			m = val
		}
	}
	return m
}

func Min[T Ordered](v1 T, v ...T) T {
	m := v1
	for _, val := range v {
		if val < m {
			m = val
		}
	}
	return m
}

func Compare[T Ordered](v1 T, v2 T) Order {
	if v1 == v2 {
		return EQ
	} else if v1 < v2 {
		return LT
	} else {
		return GT
	}
}

func Not(v bool) bool {
	return !v
}

func And(v1, v2 bool) bool {
	return v1 && v2
}

func Or(v1, v2 bool) bool {
	return v1 || v2
}

func Xor(v1, v2 bool) bool {
	return v1 != v2
}

func AppendStrings(v1, v2 string) string {
	return v1 + v2
}

func ModBy[T Int](by, base T) T {
	answer := base % by
	if (answer > 0 && by < 0) || answer < 0 && by > 0 {
		return answer + by
	}
	return answer
}

func RemainderBy[T Int](by, base T) T {
	return base % by
}

func Negate[T SignedNumbers](n T) T {
	return n * (T)(-1)
}

func Clamp[T Number](low, n, high T) T {
	if n < low {
		return low
	}
	if n > high {
		return high
	}
	return n
}

func LogBase[T Number](base, value T) T {
	return T(math.Log(float64(value)) / math.Log(float64(base)))
}

func Identity[T any](value T) T { return value }

func Always[T any, U any](a T, b U) T { return a }

func ApplyL[T, U any](f func(value T) U, val T) U {
	return f(val)
}

func ApplyR[T, U any](val T, f func(value T) U) U {
	return f(val)
}

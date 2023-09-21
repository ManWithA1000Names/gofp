package Bitwise

import "github.com/manwitha1000names/yggdrasil/fp/Basics"

func And[T Basics.Int](a, b T) T {
	return a & b
}

func Or[T Basics.Int](a, b T) T {
	return a | b
}

func Xor[T Basics.Int](a, b T) T {
	return a ^ b
}

func Not[T Basics.Int](a T) T {
	return ^a
}

func ShiftLeftBy[T Basics.Int](n T, base T) T {
	return base << n
}

func ShiftRightBy[T Basics.Int](n T, base T) T {
	return base >> n
}

package Bitwise

import "github.com/manwitha1000names/gofp/v2/Basics"

// Bitwise AND
func And[T Basics.Int](a, b T) T {
	return a & b
}

// Bitwise OR
func Or[T Basics.Int](a, b T) T {
	return a | b
}

// Bitwise XOR
func Xor[T Basics.Int](a, b T) T {
	return a ^ b
}

// Flip each bit individually, often called bitwise NOT
func Not[T Basics.Int](a T) T {
	return ^a
}

// Shift bits to the left by a given offset, filling new bits with zeros.
// This can be used to multiply numbers by powers of two.
func ShiftLeftBy[T Basics.Int](n T, base T) T {
	return base << n
}

// Shift bits to the right by a given offset, filling new bits with whatever is the topmost bit.
// This can be used to divide numbers by powers of two.
func ShiftRightBy[T Basics.Int](n T, base T) T {
	return base >> n
}

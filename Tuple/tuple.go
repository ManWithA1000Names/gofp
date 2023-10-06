package Tuple

type Tuple[T, U any] struct {
	Fst T
	Snd U
}

// Create a 2-tuple.
func Pair[T, U any](a T, b U) Tuple[T, U] {
	return Tuple[T, U]{a, b}
}

// Extract the first value from a tuple.
func First[T, U any](t Tuple[T, U]) T {
	return t.Fst
}

// Extract the second value from a tuple.
func Second[T, U any](t Tuple[T, U]) U {
	return t.Snd
}

// Transform the first value in a tuple.
func MapFirst[T, U, E any](mapfn func(value T) E, t Tuple[T, U]) Tuple[E, U] {
	return Tuple[E, U]{mapfn(t.Fst), t.Snd}
}

// Transform the second value in a tuple.
func MapSecond[T, U, E any](mapfn func(value U) E, t Tuple[T, U]) Tuple[T, E] {
	return Tuple[T, E]{t.Fst, mapfn(t.Snd)}
}

// Transform both parts of a tuple.
func MapBoth[T, U, E, D any](mapfnF func(valueF T) E, mapfnS func(valueS U) D, t Tuple[T, U]) Tuple[E, D] {
	return Tuple[E, D]{mapfnF(t.Fst), mapfnS(t.Snd)}
}

package Tuple

type Tuple[T, U any] struct {
	Fst T
	Snd U
}

func Pair[T, U any](a T, b U) Tuple[T, U] {
	return Tuple[T, U]{a, b}
}

func First[T, U any](t Tuple[T, U]) T {
	return t.Fst
}

func Second[T, U any](t Tuple[T, U]) U {
	return t.Snd
}

func MapFirst[T, U, E any](mapfn func(value T) E, t Tuple[T, U]) Tuple[E, U] {
	return Tuple[E, U]{mapfn(t.Fst), t.Snd}
}

func MapSecond[T, U, E any](mapfn func(value U) E, t Tuple[T, U]) Tuple[T, E] {
	return Tuple[T, E]{t.Fst, mapfn(t.Snd)}
}

func MapBoth[T, U, E, D any](mapfnF func(valueF T) E, mapfnS func(valueS U) D, t Tuple[T, U]) Tuple[E, D] {
	return Tuple[E, D]{mapfnF(t.Fst), mapfnS(t.Snd)}
}

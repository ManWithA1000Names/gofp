package Maybe

type Maybe[T any] struct {
	just *T
}

// CONSTRUCTION

func Just[T any](value T) Maybe[T] {
	return Maybe[T]{just: &value}
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{just: nil}
}

// METHODS

func (m *Maybe[T]) IsJust() bool {
	return m.just != nil
}

func (m *Maybe[T]) IsNothing() bool {
	return m.just == nil
}

func (m *Maybe[T]) Unwrap() T {
	return *m.just
}

// OPERATIONS

func WithDefault[T any](def T, maybe Maybe[T]) T {
	if maybe.IsJust() {
		return *maybe.just
	}
	return def
}

func Map[T, U any](mapfn func(value T) U, maybe Maybe[T]) Maybe[U] {
	if maybe.IsNothing() {
		return Nothing[U]()
	}
	return Just(mapfn(maybe.Unwrap()))
}

func Map2[A, B, C any](mapfn func(a A, b B) C, maybeA Maybe[A], maybeB Maybe[B]) Maybe[C] {
	if maybeA.IsNothing() || maybeB.IsNothing() {
		return Nothing[C]()
	}
	return Just(mapfn(maybeA.Unwrap(), maybeB.Unwrap()))
}

func Map3[A, B, C, D any](mapfn func(a A, b B, c C) D, maybeA Maybe[A], maybeB Maybe[B], maybeC Maybe[C]) Maybe[D] {
	if maybeA.IsNothing() || maybeB.IsNothing() {
		return Nothing[D]()
	}
	return Just(mapfn(maybeA.Unwrap(), maybeB.Unwrap(), maybeC.Unwrap()))
}

func Map4[A, B, C, D, E any](mapfn func(a A, b B, c C, d D) E, maybeA Maybe[A], maybeB Maybe[B], maybeC Maybe[C], maybeD Maybe[D]) Maybe[E] {
	if maybeA.IsNothing() || maybeB.IsNothing() {
		return Nothing[E]()
	}
	return Just(mapfn(maybeA.Unwrap(), maybeB.Unwrap(), maybeC.Unwrap(), maybeD.Unwrap()))
}

func Map5[A, B, C, D, E, F any](mapfn func(a A, b B, c C, d D, e E) F, maybeA Maybe[A], maybeB Maybe[B], maybeC Maybe[C], maybeD Maybe[D], maybeE Maybe[E]) Maybe[F] {
	if maybeA.IsNothing() || maybeB.IsNothing() {
		return Nothing[F]()
	}
	return Just(mapfn(maybeA.Unwrap(), maybeB.Unwrap(), maybeC.Unwrap(), maybeD.Unwrap(), maybeE.Unwrap()))
}

func AndThen[T, U any](f func(value T) Maybe[U], maybe Maybe[T]) Maybe[U] {
	if maybe.IsNothing() {
		return Nothing[U]()
	}
	return f(maybe.Unwrap())
}

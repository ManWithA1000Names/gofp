package Maybe

// Represent values that may or may not exist.
// It can be useful if you have a record field that is only filled in sometimes.
// Or if a function takes a value sometimes, but does not absolutely need it.
type Maybe[T any] struct {
	just *T
}

// CONSTRUCTION

// Create the 'Just' variant of the Maybe type.
func Just[T any](value T) Maybe[T] {
	return Maybe[T]{just: &value}
}

// Create the 'Nothing' variant of the Maybe type.
func Nothing[T any]() Maybe[T] {
	return Maybe[T]{just: nil}
}

// METHODS

// Detect wherether the Maybe is the 'Just' variant.
func (m *Maybe[T]) IsJust() bool {
	return m.just != nil
}

// Detect wherether the Maybe is the 'Nothing' variant.
func (m *Maybe[T]) IsNothing() bool {
	return m.just == nil
}

// Unwrap the Maybe type and the get the underlying value.
// It panics if Maybe is the 'Nothing' variant.
func (m *Maybe[T]) Unwrap() T {
	return *m.just
}

// OPERATIONS

// Provide a default value, turning an optional value into a normal value.
func WithDefault[T any](def T, maybe Maybe[T]) T {
	if maybe.IsJust() {
		return *maybe.just
	}
	return def
}

// Transform a Maybe value with a given function.
func Map[T, U any](mapfn func(value T) U, maybe Maybe[T]) Maybe[U] {
	if maybe.IsNothing() {
		return Nothing[U]()
	}
	return Just(mapfn(maybe.Unwrap()))
}

// Apply a function if all the arguments are Just a value.
func Map2[A, B, C any](mapfn func(a A, b B) C, maybeA Maybe[A], maybeB Maybe[B]) Maybe[C] {
	if maybeA.IsNothing() || maybeB.IsNothing() {
		return Nothing[C]()
	}
	return Just(mapfn(maybeA.Unwrap(), maybeB.Unwrap()))
}

// Apply a function if all the arguments are Just a value.
func Map3[A, B, C, D any](mapfn func(a A, b B, c C) D, maybeA Maybe[A], maybeB Maybe[B], maybeC Maybe[C]) Maybe[D] {
	if maybeA.IsNothing() || maybeB.IsNothing() {
		return Nothing[D]()
	}
	return Just(mapfn(maybeA.Unwrap(), maybeB.Unwrap(), maybeC.Unwrap()))
}

// Apply a function if all the arguments are Just a value.
func Map4[A, B, C, D, E any](mapfn func(a A, b B, c C, d D) E, maybeA Maybe[A], maybeB Maybe[B], maybeC Maybe[C], maybeD Maybe[D]) Maybe[E] {
	if maybeA.IsNothing() || maybeB.IsNothing() {
		return Nothing[E]()
	}
	return Just(mapfn(maybeA.Unwrap(), maybeB.Unwrap(), maybeC.Unwrap(), maybeD.Unwrap()))
}

// Apply a function if all the arguments are Just a value.
func Map5[A, B, C, D, E, F any](mapfn func(a A, b B, c C, d D, e E) F, maybeA Maybe[A], maybeB Maybe[B], maybeC Maybe[C], maybeD Maybe[D], maybeE Maybe[E]) Maybe[F] {
	if maybeA.IsNothing() || maybeB.IsNothing() {
		return Nothing[F]()
	}
	return Just(mapfn(maybeA.Unwrap(), maybeB.Unwrap(), maybeC.Unwrap(), maybeD.Unwrap(), maybeE.Unwrap()))
}

// CHAINING METHODS

// Chain together many computations that may fail.
func AndThen[T, U any](f func(value T) Maybe[U], maybe Maybe[T]) Maybe[U] {
	if maybe.IsNothing() {
		return Nothing[U]()
	}
	return f(maybe.Unwrap())
}

func Or[T any](m Maybe[T], maybe Maybe[T]) Maybe[T] {
	if maybe.IsJust() {
		return maybe
	}
	return m
}

func And[T any](m Maybe[T], maybe Maybe[T]) Maybe[T] {
	if maybe.IsJust() {
		return m
	}
	return maybe
}

package MaybeResult

import (
	"fmt"
)

var unwrapped_nothing = fmt.Errorf("Tried to unwrap a `Maybe` that was `Nothing`.")

// Represent values that may or may not exist.
// It can be useful if you have a record field that is only filled in sometimes.
// Or if a function takes a value sometimes, but does not absolutely need it.
type Maybe[T any] struct {
	isJust bool
	value  T
}

// CONSTRUCTION

// Create a `Maybe` from a function that returns (T, bool)
func TupleToMaybe[T any](value T, ok bool) Maybe[T] {
	if ok {
		return Just(value)
	}
	return Nothing[T]()
}

// Create a `Maybe` from a function that returns (T, error)
func ErrToMaybe[T any](value T, err error) Maybe[T] {
	if err != nil {
		return Nothing[T]()
	}
	return Just(value)
}

// Create a `Maybe` from a possibly nil pointer
func PtrToMaybe[T any](value *T) Maybe[T] {
	if value == nil {
		return Nothing[T]()
	}
	return Just(*value)
}

// Create the 'Just' variant of the Maybe type.
func Just[T any](value T) Maybe[T] {
	return Maybe[T]{isJust: true, value: value}
}

// Create the 'Nothing' variant of the Maybe type.
func Nothing[T any]() Maybe[T] {
	return Maybe[T]{isJust: false}
}

// METHODS

func (m Maybe[T]) ToPtr() *T {
	if m.isJust {
		return &m.value
	}

	return nil
}

// Detect wherether the Maybe is the 'Just' variant.
func (m Maybe[T]) IsJust() bool {
	return m.isJust
}

// Detect wherether the Maybe is the 'Nothing' variant.
func (m Maybe[T]) IsNothing() bool {
	return !m.isJust
}

// Unwrap the `Maybe` type and the get the underlying value.
//
//	Just(T) => T
//	Nothing => PANIC!
func (m Maybe[T]) Unwrap() T {
	if !m.isJust {
		panic(unwrapped_nothing)
	}
	return m.value
}

// Unwrap the `Maybe` type and get the underlying value, or the default value of the type.
//
//	Just(T) => T
//	Nothing => default[T]()
func (m Maybe[T]) OrDefault() T {
	return m.value
}

// Unwrap the `Maybe` type and get the underlying value, or the value provided.
//
//	Just(T) => T
//	Nothing => parameter value T
func (m Maybe[T]) OrValue(value T) T {
	if m.isJust {
		return m.value
	}
	return value
}

// Unwrap the `Maybe` type and get the underlying value, or execute the provided fn.
//
//	Just(T) => T
//	Nothing => fn()
func (m Maybe[T]) OrElse(fn func() T) T {
	if m.isJust {
		return m.value
	}
	return fn()
}

// Execute the fn on the underlying value if the `Maybe` is Just
// returning a `Maybe` with the returned value.
//
//	Just(T) => Just(fn(T))
//	Nothing => Nothing
func (m Maybe[T]) Map(fn func(value T) T) Maybe[T] {
	if m.isJust {
		return Maybe[T]{
			isJust: true,
			value:  fn(m.value),
		}
	}
	return m
}

// Chain together many computations that may fail.
//
//	Just(T) => fn(T)
//	Nothing => Nothing
func (m Maybe[T]) AndThen(fn func(value T) Maybe[T]) Maybe[T] {
	if m.isJust {
		return fn(m.value)
	}

	return m
}

func (m Maybe[T]) Then(fn func(value T)) {
	if m.isJust {
		fn(m.value)
	}
}

// Return the other Maybe if this one is Nothing.
//
//	Just => this
//	Nothing => other
func (m Maybe[T]) Or(other Maybe[T]) Maybe[T] {
	if m.isJust {
		return m
	}

	return other
}

// Return the other Maybe if this is Just.
//
//	Just => other
//	Nothing => this
func (m Maybe[T]) And(other Maybe[T]) Maybe[T] {
	if m.isJust {
		return other
	}

	return m
}

func (m Maybe[T]) OkOr(err error) Result[T] {
	if m.isJust {
		return Ok(m.value)
	}
	return Err[T](err)
}

func (m Maybe[T]) ToResult() Result[T] {
	return m.OkOr(fmt.Errorf("Maybe was Nothing."))
}

// INTERFACE IMPLEMENTATIONS

func (m Maybe[T]) Format(f fmt.State, c rune) {
	if m.isJust {
		f.Write([]byte("Just(" + fmt.Sprint(m.value) + ")"))
	} else {
		f.Write([]byte("Nothing"))
	}
}

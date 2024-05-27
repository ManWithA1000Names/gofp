package MaybeResult

import (
	"fmt"
)

var unwrapping_err_while_is_ok = fmt.Errorf("Tried UnrwappingErr on a Ok value.")

// A Result is either Ok meaning the computation succeeded,
// or it is an Err meaning that there was some failure.
type Result[T any] struct {
	err   error
	value T
}

// CREATION

func ErrToResult[T any](value T, err error) Result[T] {
	return Result[T]{
		err:   err,
		value: value,
	}
}

func Try[T any](fn func() (T, error)) Result[T] {
	return ErrToResult(fn())
}

// Create the 'Ok' variant of the Result type.
func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

// Create the 'Err' variant of the Result type.
func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func Errf[T any](format string, a ...any) Result[T] {
	return Err[T](fmt.Errorf(format, a...))
}

// METHODS

// Detect wherether the Result is the 'Ok' variant.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// Detect wherether the Result is the 'Ok' variant.
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Unwrap the Result type and the get the underlying 'Ok' value.
// It panics if Result is the 'Err' variant.
func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(r.err)
	}
	return r.value
}

// Unwrap the Result type and the get the underlying 'Err' value.
// It panics if Result is the 'Ok' variant.
func (r Result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic(unwrapping_err_while_is_ok)
	}
	return r.err
}

func (r Result[T]) Ok() Maybe[T] {
	if r.IsOk() {
		return Just(r.value)
	}
	return Nothing[T]()
}

func (r Result[T]) Err() Maybe[error] {
	if r.IsOk() {
		return Nothing[error]()
	}
	return Just(r.err)
}

func (r Result[T]) OrDefault() T {
	return r.value
}

func (r Result[T]) OrValue(value T) T {
	if r.IsOk() {
		return r.value
	}
	return value
}

func (r Result[T]) OrElse(fn func() T) T {
	if r.IsOk() {
		return r.value
	}
	return fn()
}

func (r Result[T]) Map(fn func(value T) T) Result[T] {
	if r.IsOk() {
		return Result[T]{
			value: fn(r.value),
		}
	}
	return r
}

func (r Result[T]) MapErr(fn func(err error) error) Result[T] {
	if r.IsOk() {
		return r
	}
	return Result[T]{
		err: fn(r.err),
	}
}

func (r Result[T]) AndThen(fn func(value T) Result[T]) Result[T] {
	if r.IsOk() {
		return fn(r.value)
	}
	return r
}

func (r Result[T]) OrThen(fn func(err error) Result[T]) Result[T] {
	if r.IsErr() {
		return fn(r.err)
	}
	return r
}

func (r Result[T]) Or(other Result[T]) Result[T] {
	if r.IsOk() {
		return r
	}
	return other
}

func (r Result[T]) And(other Result[T]) Result[T] {
	if r.IsOk() {
		return other
	}
	return r
}

func (r Result[T]) ToMaybe() Maybe[T] {
	return r.Ok()
}

// INTERFACE IMPLEMENTATIONS

func (r Result[T]) Format(f fmt.State, c rune) {
	if r.IsOk() {
		f.Write([]byte("Ok(" + fmt.Sprint(r.value) + ")"))
	} else {
		f.Write([]byte("Err(" + r.err.Error() + ")"))
	}
}

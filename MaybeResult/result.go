package MaybeResult

import (
	"errors"
	"fmt"
)

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

// Expects the Result to b e of the `Ok` variant and gets the underlying value.
// It panics if Result is the 'Err' variant.
func (r Result[T]) Expect() T {
	if r.IsErr() {
		panic(r.err)
	}
	return r.value
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

func (r Result[T]) WithDefault(value T) T {
	if r.IsOk() {
		return r.value
	}
	return value
}

func (r Result[T]) OrElse(fn func(err error) T) T {
	if r.IsOk() {
		return r.value
	}
	return fn(r.err)
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
		_, _ = f.Write([]byte("Ok(" + fmt.Sprint(r.value) + ")"))
	} else {
		_, _ = f.Write([]byte("Err(" + r.err.Error() + ")"))
	}
}

// For working with the `errors`.Unwrap functionallity.
func (r Result[T]) Unwrap() error {
	if r.IsOk() {
		return nil
	}
	return r.err
}

// Result implenents it self the `error` interface, proxying for its underlying error if there is one.
func (r Result[T]) Error() string {
	err := r.Unwrap()
	if err == nil {
		panic(fmt.Errorf("Called the `Error` method on a `Result` of the `Ok` variant."))
	}

	return err.Error()
}

// For the use with the errors.As function.
func (r Result[T]) As(target any) bool {
	if r.IsOk() {
		return false
	}
	return errors.As(r.err, target)
}

// For the use with the errors.Is function.
func (r Result[T]) Is(err error) bool {
	if r.IsOk() {
		return false
	}
	return errors.Is(r.err, err)
}

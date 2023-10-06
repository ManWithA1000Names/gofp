package Result

import (
	"fmt"

	"github.com/manwitha1000names/gofp/Maybe"
)

// A Result is either Ok meaning the computation succeeded,
// or it is an Err meaning that there was some failure.
type Result[T any] struct {
	err error
	ok  *T
}

// CREATION

// Create the 'Ok' variant of the Result type.
func Ok[T any](value T) Result[T] {
	return Result[T]{ok: &value, err: nil}
}

// Create the 'Err' variant of the Result type.
func Err[T any](err error) Result[T] {
	return Result[T]{ok: nil, err: err}
}

// METHODS

// Detect wherether the Result is the 'Ok' variant.
func (r *Result[T]) IsOk() bool {
	return r.ok != nil
}

// Detect wherether the Result is the 'Ok' variant.
func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

// Unwrap the Result type and the get the underlying 'Ok' value.
// It panics if Result is the 'Err' variant.
func (r *Result[T]) Unwrap() T {
	return *r.ok
}

// Unwrap the Result type and the get the underlying 'Err' value.
// It panics if Result is the 'Ok' variant.
func (r *Result[T]) UnwrapErr() error {
	return r.err
}

// UTILITIES

// If the result is Ok return the value,
// but if the result is an Err then return a given default value.
func WithDefault[T any](def T, result Result[T]) T {
	if result.IsErr() {
		return def
	}
	return result.Unwrap()
}

// Convert to a simpler Maybe if the actual error message is not needed or you need to interact with some code that primarily uses maybes.
func ToMaybe[T any](result Result[T]) Maybe.Maybe[T] {
	if result.IsErr() {
		return Maybe.Nothing[T]()
	}
	return Maybe.Just(result.Unwrap())
}

// Convert from a simple Maybe to interact with some code that primarily uses Results.
func FromMaybe[T any](err error, maybe Maybe.Maybe[T]) Result[T] {
	if maybe.IsJust() {
		return Ok(maybe.Unwrap())
	}
	return Err[T](err)
}

// Transform an Err value.
func MapError[T any](mapfn func(err error) error, result Result[T]) Result[T] {
	if result.IsOk() {
		return Ok(result.Unwrap())
	}
	return Err[T](mapfn(result.UnwrapErr()))
}

// Chain together a sequence of computations that may fail.
func AndThen[T, U any](f func(value T) Result[U], result Result[T]) Result[U] {
	if result.IsErr() {
		return Err[U](result.UnwrapErr())
	}
	return f(result.Unwrap())
}

// Get  the 'Ok' value from the first Result,
// else try to get the 'Ok' from the second result,
// else 'Err' is result.
func Or[T any](r Result[T], result Result[T]) Result[T] {
	if result.IsErr() {
		return r
	}
	return result
}

// If the first result is 'Ok' then get the second result,
// else get the first 'Err'.
func And[T any](r Result[T], result Result[T]) Result[T] {
	if result.IsErr() {
		return result
	}
	return r
}

// Apply a function to a result. If the result is Ok, it will be converted.
// If the result is an Err, the same error value will propagate through.
func Map[T, U any](mapfn func(value T) U, result Result[T]) Result[U] {
	if result.IsErr() {
		return Err[U](result.UnwrapErr())
	}
	return Ok(mapfn(result.Unwrap()))
}

// Apply a function if both results are Ok. If not, the first Err will propagate through.
func Map2[a, b, value any](mapfn func(a a, b b) value, resulta Result[a], resultb Result[b]) Result[value] {
	if resulta.IsErr() {
		return Err[value](resulta.UnwrapErr())
	}
	if resultb.IsErr() {
		return Err[value](resultb.UnwrapErr())
	}
	return Ok(mapfn(resulta.Unwrap(), resultb.Unwrap()))
}

// Apply a function if all results are Ok. If not, the first Err will propagate through.
func Map3[a, b, c, value any](mapfn func(a a, b b, c c) value, resulta Result[a], resultb Result[b], resultc Result[c]) Result[value] {
	if resulta.IsErr() {
		return Err[value](resulta.UnwrapErr())
	}
	if resultb.IsErr() {
		return Err[value](resultb.UnwrapErr())
	}
	if resultc.IsErr() {
		return Err[value](resultc.UnwrapErr())
	}
	return Ok(mapfn(resulta.Unwrap(), resultb.Unwrap(), resultc.Unwrap()))
}

// Apply a function if all results are Ok. If not, the first Err will propagate through.
func Map4[a, b, c, d, value any](mapfn func(a a, b b, c c, d d) value, resulta Result[a], resultb Result[b], resultc Result[c], resultd Result[d]) Result[value] {
	if resulta.IsErr() {
		return Err[value](resulta.UnwrapErr())
	}
	if resultb.IsErr() {
		return Err[value](resultb.UnwrapErr())
	}
	if resultc.IsErr() {
		return Err[value](resultc.UnwrapErr())
	}
	if resultd.IsErr() {
		return Err[value](resultd.UnwrapErr())
	}
	return Ok(mapfn(resulta.Unwrap(), resultb.Unwrap(), resultc.Unwrap(), resultd.Unwrap()))
}

// Apply a function if all results are Ok. If not, the first Err will propagate through.
func Map5[a, b, c, d, e, value any](mapfn func(a a, b b, c c, d d, e e) value, resulta Result[a], resultb Result[b], resultc Result[c], resultd Result[d], resulte Result[e]) Result[value] {
	if resulta.IsErr() {
		return Err[value](resulta.UnwrapErr())
	}
	if resultb.IsErr() {
		return Err[value](resultb.UnwrapErr())
	}
	if resultc.IsErr() {
		return Err[value](resultc.UnwrapErr())
	}
	if resultd.IsErr() {
		return Err[value](resultd.UnwrapErr())
	}
	if resulte.IsErr() {
		return Err[value](resulte.UnwrapErr())
	}
	return Ok(mapfn(resulta.Unwrap(), resultb.Unwrap(), resultc.Unwrap(), resultd.Unwrap(), resulte.Unwrap()))
}

// GO SPECIFIC

// Create a Result from the callic: value, ok := fn().
//
// Just use it like so: result := Result.FromValueOk(fn())
func FromValueOk[T any](value T, ok bool) Result[T] {
	if ok {
		return Err[T](fmt.Errorf("Operation failed."))
	}
	return Ok(value)
}

// Same as FromValueOk but the value is a pointer.
func FromPtrValueOk[T any](value *T, ok bool) Result[T] {
	if ok {
		return Err[T](fmt.Errorf("Operation failed."))
	}
	return Result[T]{ok: value, err: nil}
}

// Create a Result from the callic: value, err := fn().
//
// Just use it like so: result := Result.FromValueOk(fn())
func FromValueErr[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(value)
}

// Same as FromValueErr but the value is a pointer.
func FromPtrValueErr[T any](value *T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Result[T]{ok: value, err: nil}
}

package Result

import "github.com/manwitha1000names/yggdrasil/fp/Maybe"

type Result[T any] struct {
	err error
	ok  *T
}

// CREATION

func Ok[T any](value T) Result[T] {
	return Result[T]{ok: &value, err: nil}
}

func Err[T any](err error) Result[T] {
	return Result[T]{ok: nil, err: err}
}

// METHODS

func (r *Result[T]) IsOk() bool {
	return r.ok != nil
}

func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

func (r *Result[T]) Unwrap() T {
	return *r.ok
}

func (r *Result[T]) UnwrapErr() error {
	return r.err
}

// UTILITIES

func WithDefault[T any](def T, result Result[T]) T {
	if result.IsErr() {
		return def
	}
	return result.Unwrap()
}

func ToMaybe[T any](result Result[T]) Maybe.Maybe[T] {
	if result.IsErr() {
		return Maybe.Nothing[T]()
	}
	return Maybe.Just(result.Unwrap())
}

func FromMaybe[T any](err error, maybe Maybe.Maybe[T]) Result[T] {
	if maybe.IsJust() {
		return Ok(maybe.Unwrap())
	}
	return Err[T](err)
}

func MapError[T any](mapfn func(err error) error, result Result[T]) Result[T] {
	if result.IsOk() {
		return Ok(result.Unwrap())
	}
	return Err[T](mapfn(result.UnwrapErr()))
}

func AndThen[T, U any](f func(value T) Result[U], result Result[T]) Result[U] {
	if result.IsErr() {
		return Err[U](result.UnwrapErr())
	}
	return f(result.Unwrap())
}

func Map[T, U any](mapfn func(value T) U, result Result[T]) Result[U] {
	if result.IsErr() {
		return Err[U](result.UnwrapErr())
	}
	return Ok(mapfn(result.Unwrap()))
}

func Map2[a, b, value any](mapfn func(a a, b b) value, resulta Result[a], resultb Result[b]) Result[value] {
	if resulta.IsErr() {
		return Err[value](resulta.UnwrapErr())
	}
	if resultb.IsErr() {
		return Err[value](resultb.UnwrapErr())
	}
	return Ok(mapfn(resulta.Unwrap(), resultb.Unwrap()))
}

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

package Result

import "fmt"

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

// Create `Result` from a possibly nil pointer
func FromPtr[T any](value *T) Result[T] {
	if value == nil {
		return Err[T](fmt.Errorf("Pointer is nil"))
	}
	return Result[T]{ok: value, err: nil}
}

package Maybe

// GO SPECIFIC

// Create a `Maybe` from a function that returns (T, bool)
func FromValueOk[T any](value T, ok bool) Maybe[T] {
	if ok {
		return Just(value)
	}
	return Nothing[T]()
}

// Create a `Maybe` from a function that returns (*T, bool)
func FromValuePtrOk[T any](value *T, ok bool) Maybe[T] {
	if ok {
		return Maybe[T]{just: value}
	}
	return Nothing[T]()
}

// Create a `Maybe` from a function that returns (T, error)
func FromValueErr[T any](value T, err error) Maybe[T] {
	if err != nil {
		return Nothing[T]()
	}
	return Just(value)
}

// Create a `Maybe` from a function that returns (*T, error)
func FromPtrValueErr[T any](value *T, err error) Maybe[T] {
	if err != nil {
		return Nothing[T]()
	}
	return Maybe[T]{just: value}
}

// Create a `Maybe` from a possibly nil pointer
func FromPtr[T any](value *T) Maybe[T] {
	return Maybe[T]{just: value}
}

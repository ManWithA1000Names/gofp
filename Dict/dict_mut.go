package Dict

// BUILD

// Insert a key-value pair into a dictionary. Replaces value when there is a collision.
// This functions is MUTABLE and will change the dict in place.
func Insert_mut[Key comparable, Value any](key Key, v Value, m map[Key]Value) map[Key]Value {
	m[key] = v
	return m
}

// Update the value of a dictionary for a specific key with a given function.
// This functions is MUTABLE and will change the dict in place.
func Update_mut[Key comparable, Value any](key Key, upfn func(value Value) Value, m map[Key]Value) map[Key]Value {
	if value, ok := m[key]; ok {
		m[key] = upfn(value)
	}
	return m
}

// Remove a key-value pair from a dictionary. If the key is not found, no changes are made.
// This functions is MUTABLE and will change the dict in place.
func Remove_mut[Key comparable, Value any](key Key, m map[Key]Value) map[Key]Value {
	delete(m, key)
	return m
}

// TRANSFORM

// Apply a function to all values in a dictionary.
// This functions is MUTABLE and will change the dict in place.
func Map_mut[Key comparable, Value any](mapfn func(key Key, value Value) Value, m map[Key]Value) map[Key]Value {
	for _, key := range Keys(m) {
		m[key] = mapfn(key, m[key])
	}
	return m
}

// Keep only the key-value pairs that pass the given test.
// This functions is MUTABLE and will change the dict in place.
func Filter_mut[Key comparable, Value any](testfn func(key Key, value Value) bool, m map[Key]Value) map[Key]Value {
	for key, value := range m {
		if !testfn(key, value) {
			delete(m, key)
		}
	}
	return m
}

// COMBINE

// Combine two dictionaries. If there is a collision, preference is given to the first dictionary.
// This functions is MUTABLE and will change the first dict in place.
// The first dict is mutated and returned.
func Union_mut[Key comparable, Value any](m map[Key]Value, m1 map[Key]Value) map[Key]Value {
	for key, value := range m1 {
		if _, ok := m[key]; !ok {
			m[key] = value
		}
	}
	return m
}

// Keep a key-value pair when its key appears in the second dictionary. Preference is given to values in the first dictionary.
// This functions is MUTABLE and will change the first dict in place.
// The first dict is mutated and returned.
func Intersect_mut[Key comparable, Value any](m map[Key]Value, m1 map[Key]Value) map[Key]Value {
	return Filter_mut(func(key Key, _ Value) bool {
		return Member(key, m1)
	}, m)
}

// Keep a key-value pair when its key does not appear in the second dictionary.
// This functions is MUTABLE and will change the first dict in place.
// The first dict is mutated and returned.
func Diff_mut[Key comparable, Value any](m map[Key]Value, m1 map[Key]Value) map[Key]Value {
	return Filter_mut(func(key Key, _ Value) bool {
		return !Member(key, m1)
	}, m)
}

package Set

import "github.com/manwitha1000names/gofp/Dict"

// BUILD

// Insert a value into a set.
// This functions is MUTABLE and will change the set in place.
func Insert_mut[T comparable](value T, s Set[T]) Set[T] {
	s.m = Dict.Insert_mut(value, 0, s.m)
	return s
}

// Remove a value from a set.
// If the value is not found, no changes are made.
// This functions is MUTABLE and will change the set in place.
func Remove_mut[T comparable](value T, s Set[T]) Set[T] {
	s.m = Dict.Remove_mut(value, s.m)
	return s
}

// COMBINE

// Get the union of two sets. Keep all values.
// This functions is MUTABLE and will change the first set in place.
func Union_mut[T comparable](s Set[T], s1 Set[T]) Set[T] {
	s.m = Dict.Union_mut(s.m, s1.m)
	return s
}

// Get the intersection of two sets. Keeps values that appear in both sets.
// This functions is MUTABLE and will change the first set in place.
func Intersect_mut[T comparable](s Set[T], s1 Set[T]) Set[T] {
	s.m = Dict.Intersect_mut(s.m, s1.m)
	return s
}

// Get the difference between the first set and the second.
// Keeps values that do not appear in the second set.
// This functions is MUTABLE and will change the first set in place.
func Diff_mut[T comparable](s Set[T], s1 Set[T]) Set[T] {
	s.m = Dict.Diff_mut(s.m, s1.m)
	return s
}

// TRANSFORM

// Map a function onto a set, creating a new set with no duplicates.
// This functions is MUTABLE and will change the set in place.
func Map_mut[T comparable](mapfn func(value T) T, s Set[T]) Set[T] {
	for key := range s.m {
		s.m[mapfn(key)] = 0
	}
	return s
}

// Only keep elements that pass the given test.
// This functions is MUTABLE and will change the set in place.
func Filter_mut[T comparable](testfn func(value T) bool, s Set[T]) Set[T] {
	s.m = Dict.Filter_mut(func(value T, _ int) bool {
		return testfn(value)
	}, s.m)
	return s
}

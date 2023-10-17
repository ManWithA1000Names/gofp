package Set

import (
	"github.com/manwitha1000names/gofp/Dict"
	"github.com/manwitha1000names/gofp/List"
	"github.com/manwitha1000names/gofp/Tuple"
)

// Represents a set of unique values.
type Set[T comparable] struct {
	m map[T]struct{}
}

// BUILD

// Create an empty set.
func Empty[T comparable]() Set[T] {
	return Set[T]{make(map[T]struct{})}
}

// Create a set with one value.
func Singleton[T comparable](value T) Set[T] {
	new_map := make(map[T]struct{})
	new_map[value] = struct{}{}
	return Set[T]{new_map}
}

// Insert a value into a set.
// This functions is IMMUTABLE and produces a completely new set!
func Insert[T comparable](value T, s Set[T]) Set[T] {
	return Set[T]{Dict.Insert(value, struct{}{}, s.m)}
}

// Remove a value from a set.
// If the value is not found, no changes are made.
// This functions is IMMUTABLE and produces a completely new set!
func Remove[T comparable](value T, s Set[T]) Set[T] {
	return Set[T]{Dict.Remove(value, s.m)}
}

// QUERY

// Determine if a set is empty.
func IsEmpty[T comparable](s Set[T]) bool {
	return Dict.IsEmpty(s.m)
}

// Determine if a value is in a set.
func Member[T comparable](value T, s Set[T]) bool {
	return Dict.Member(value, s.m)
}

// Determine the number of elements in a set.
func Size[T comparable](s Set[T]) int {
	return len(s.m)
}

// COMBINE

// Get the union of two sets. Keep all values.
// This functions is IMMUTABLE and produces a completely new set!
func Union[T comparable](s Set[T], s1 Set[T]) Set[T] {
	return Set[T]{Dict.Union(s.m, s1.m)}
}

// Get the intersection of two sets. Keeps values that appear in both sets.
// This functions is IMMUTABLE and produces a completely new set!
func Intersect[T comparable](s Set[T], s1 Set[T]) Set[T] {
	return Set[T]{Dict.Intersect(s.m, s1.m)}
}

// Get the difference between the first set and the second.
// Keeps values that do not appear in the second set.
// This functions is IMMUTABLE and produces a completely new set!
func Diff[T comparable](s Set[T], s1 Set[T]) Set[T] {
	return Set[T]{Dict.Diff(s.m, s1.m)}
}

// LISTS

// Convert a set into a list, sorted from lowest to highest.
func ToList[T comparable](s Set[T]) []T {
	return Dict.Keys(s.m)
}

// Convert a list into a set, removing any duplicates.
func FromList[T comparable](list []T) Set[T] {
	return Set[T]{Dict.FromList(List.Map(func(value T) Tuple.Tuple[T, struct{}] {
		return Tuple.Pair(value, struct{}{})
	}, list))}
}

// TRANSFORM

// Map a function onto a set, creating a new set with no duplicates.
// This functions is IMMUTABLE and produces a completely new set!
func Map[T comparable, U comparable](mapfn func(value T) U, s Set[T]) Set[U] {
	new_map := make(map[U]struct{}, len(s.m))
	for key := range s.m {
		new_map[mapfn(key)] = struct{}{}
	}
	return Set[U]{new_map}
}

// Fold over the values in a set, in order from lowest to highest.
// This functions is IMMUTABLE and produces a completely new set!
func Foldl[T comparable, Acc any](reducer func(value T, acc Acc) Acc, init Acc, s Set[T]) Acc {
	return List.Foldl(reducer, init, ToList(s))
}

// Fold over the values in a set, in order from highest to lowest.
// This functions is IMMUTABLE and produces a completely new set!
func Foldr[T comparable, Acc any](reducer func(value T, acc Acc) Acc, init Acc, s Set[T]) Acc {
	return List.Foldr(reducer, init, ToList(s))
}

// Only keep elements that pass the given test.
// This functions is IMMUTABLE and produces a completely new set!
func Filter[T comparable](testfn func(value T) bool, s Set[T]) Set[T] {
	return Set[T]{Dict.Filter(func(value T, _ struct{}) bool {
		return testfn(value)
	}, s.m)}
}

// Create two new sets.
// The first contains all the elements that passed the given test,
// and the second contains all the elements that did not.
func Partition[T comparable](partfn func(value T) bool, s Set[T]) (Set[T], Set[T]) {
	lista, listb := List.Partition(partfn, ToList(s))
	return FromList(lista), FromList(listb)
}

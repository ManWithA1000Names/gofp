package List

import (
	"math/bits"

	"github.com/manwitha1000names/gofp/v3/Basics"
	. "github.com/manwitha1000names/gofp/v3/MaybeResult"
)

// TRANSFORM

// Apply a function to every element of a list.
// This functions is MUTABLE and will change the list in place.
func Map_mut[T any](mapfn func(value T) T, list []T) []T {
	for i, v := range list {
		list[i] = mapfn(v)
	}
	return list
}

// Same as map but the function is also applied to the index of each element (starting at zero).
// This functions is MUTABLE and will change the list in place.
func IndexedMap_mut[T any](f func(index int, value T) T, list []T) []T {
	for i, w := range list {
		list[i] = f(i, w)
	}
	return list
}

// Keep elements that satisfy the test.
// This functions is MUTABLE and will change the list in place.
func Filter_mut[T any](testfn func(value T) bool, list []T) []T {
	i := 0
	for _, v := range list {
		if testfn(v) {
			list[i] = v
			i += 1
		}
	}
	return list[:i]
}

// Filter out certain values.
// This functions is MUTABLE and will change the list in place.
func FilterMap_mut[T any](testmapfn func(value T) Maybe[T], list []T) []T {
	i := 0
	for _, v := range list {
		new_value := testmapfn(v)
		if new_value.IsJust() {
			list[i] = new_value.Expect()
			i += 1
		}
	}
	return list[:i]
}

// Reverse a list.
// This functions is MUTABLE and will change the list in place.
func Reverse_mut[T any](list []T) []T {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

// COMBINE

// Put two lists together.
// This functions is MUTABLE and will change the first list in place.
// This will append the entire listB to listA and return listA.
func Append_mut[T any](listA []T, listB []T) []T {
	return append(listA, listB...)
}

// Concatenate a bunch of lists into a single list:
// This functions is MUTABLE and will change the first list in the list in place.
// This will append all the list[1:] to the first list and return the first list.
// If the list is empty it will return a new empty list.
func Concat_mut[T any](list [][]T) []T {
	if len(list) == 0 {
		return []T{}
	}
	return Foldl(func(value []T, acc []T) []T {
		return append(acc, value...)
	}, list[0], list[1:])
}

// Map a given function onto a list and flatten the resulting lists.
// This functions is MUTABLE and will change the first list in the list in place.
// This will append all the list[1:] to the first list and return the first list.
// If the list is empty it will return a new empty list.
func ConcatMap_mut[T any](mapfn func(value []T) []T, list [][]T) []T {
	return Concat_mut(Map_mut(mapfn, list))
}

// Places the given value between all members of the given list.
// This functions is MUTABLE and will change the list in place.
func Intersperse_mut[T any](value T, list []T) []T {
	// if the list length is less than 2
	// it is not able to be interspersed.
	if len(list) < 2 {
		return list
	}

	// We need this for later
	original_length := len(list)

	// We will add len(list) - 1 new elements to the list.
	// Thus we can now extend the length & capacity of the array to fit the new elements.
	for i := 0; i < original_length-1; i++ {
		list = append(list, value) // TODO: Since this at some point just copies the underlying array, lets just do it our selves.
	}

	// Swap all the elements to the correct place.
	// At this stage all the new elements are bunched up at the back
	// and we need to swap them to their correct location.
	// The formula for the existing elements: new_index = old_index.
	// We need to start from the back to avoid ovveriding values.
	for i := original_length - 1; i > 0; i -= 1 {
		list[i*2], list[i] = list[i], list[i*2]
	}

	return list
}

// SORT

// Sort values from lowest to highest
// This functions is MUTABLE and will change the list in place.
func Sort_mut[T Basics.Ordered](list []T) []T {
	n := len(list)
	pdqsortOrdered(list, 0, n, bits.Len(uint(n)))
	return list
}

// Sort values by a derived property.
// This functions is MUTABLE and will change the list in place.
func SortBy_mut[T any, U Basics.Ordered](mapfn func(value T) U, list []T) []T {
	n := len(list)
	pdqsortCmpFunc(list, 0, n, bits.Len(uint(n)), func(a, b T) Basics.Order {
		return Basics.Compare(mapfn(a), mapfn(b))
	})
	return list
}

// Sort values with a custom comparison function.
// This functions is MUTABLE and will change the list in place.
func SortWith_mut[T any](cmpfn func(a, b T) Basics.Order, list []T) []T {
	n := len(list)
	pdqsortCmpFunc(list, 0, n, bits.Len(uint(n)), cmpfn)
	return list
}

// FROM ARRAY

// Set the element at a particular index. Returns an updated array.
// If the index is out of range, the array is unaltered.
// This functions is MUTABLE and will change the list in place.
func Set_mut[T any](index int, value T, list []T) []T {
	if index < 0 || index >= len(list) {
		return list
	}
	list[index] = value
	return list
}

// Push an element onto the end of an array.
// This functions is MUTABLE and will change the list in place.
func Push_mut[T any](value T, list []T) []T {
	return append(list, value)
}

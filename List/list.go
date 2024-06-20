package List

import (
	"math/bits"

	"github.com/manwitha1000names/gofp/v3/Basics"
	. "github.com/manwitha1000names/gofp/v3/MaybeResult"
	"github.com/manwitha1000names/gofp/v3/Tuple"
)

func clone[T any](value []T) []T {
	new_list := make([]T, len(value))
	copy(new_list, value)
	return new_list
}

// CREATE

// Create a list with only one element
func Singleton[T any](value T) []T {
	return []T{value}
}

// Create a list with n copies of a value
func Repeat[T any](amount int, value T) []T {
	list := make([]T, 0, Basics.Max(amount, 8))
	for i := 0; i < amount; i++ {
		list = append(list, value)
	}
	return list
}

// Create a list of numbers, every element increasing by one.
// You give the lowest and highest number that should be in the list.
func Range(start, end int) []int {
	if start > end {
		return []int{}
	}
	length := end - start + 1
	list := make([]int, 0, length)
	for i := 0; i < length; i++ {
		list = append(list, i+start)
	}
	return list
}

// Add an element to the front of a list.
// This functions is IMMUTABLE and produces a completely new list!
func Cons[T any](value T, list []T) []T {
	new_list := make([]T, 0, 1+len(list))
	return append(append(new_list, value), list...)
}

// TRANSFORM

// Apply a function to every element of a list.
// This functions is IMMUTABLE and produces a completely new list!
func Map[T, U any](mapfn func(value T) U, list []T) []U {
	new_list := make([]U, 0, len(list))
	for _, w := range list {
		new_list = append(new_list, mapfn(w))
	}
	return new_list
}

// Same as map but the function is also applied to the index of each element (starting at zero).
// This functions is IMMUTABLE and produces a completely new list!
func IndexedMap[T, U any](mapfn func(index int, value T) U, list []T) []U {
	new_list := make([]U, 0, len(list))
	for i, w := range list {
		new_list = append(new_list, mapfn(i, w))
	}
	return new_list
}

// Reduce a list from the left.
func Foldl[T, U any](reducefn func(value T, accumulator U) U, acc U, list []T) U {
	for _, w := range list {
		acc = reducefn(w, acc)
	}
	return acc
}

// Reduce a list from the right.
func Foldr[T, U any](reducefn func(value T, accumulator U) U, acc U, list []T) U {
	for i := len(list) - 1; i >= 0; i-- {
		acc = reducefn(list[i], acc)
	}
	return acc
}

// Keep elements that satisfy the test.
// This functions is IMMUTABLE and produces a completely new list!
func Filter[T any](testfn func(value T) bool, list []T) []T {
	new_list := make([]T, 0, len(list))
	for _, v := range list {
		if testfn(v) {
			new_list = append(new_list, v)
		}
	}
	return new_list
}

// Filter out certain values.
// This functions is IMMUTABLE and produces a completely new list!
func FilterMap[T, U any](testmapfn func(value T) Maybe[U], list []T) []U {
	new_list := make([]U, 0, len(list))
	for _, v := range list {
		res := testmapfn(v)
		if res.IsJust() {
			new_list = append(new_list, res.Expect())
		}
	}
	return new_list
}

// UTILITIES

// Determine the length of a list.
func Length[T any](list []T) int {
	return len(list)
}

// Reverse a list.
// This functions is IMMUTABLE and produces a completely new list!
func Reverse[T any](list []T) []T {
	return Foldl(Cons[T], []T{}, list)
}

// Figure out whether a list contains a value.
func Member[T comparable](value T, list []T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// Determine if all elements satisfy some test.
func All[T any](test func(value T) bool, list []T) bool {
	for _, v := range list {
		if !test(v) {
			return false
		}
	}
	return true
}

// Determine if any elements satisfy some test.
func Any[T any](test func(value T) bool, list []T) bool {
	for _, v := range list {
		if test(v) {
			return true
		}
	}
	return false
}

// Find the maximum element in a non-empty list
func Maximum[T Basics.Ordered](list []T) Maybe[T] {
	if len(list) == 0 {
		return Nothing[T]()
	}
	return Just(Basics.Max(list[0], list[1:]...))
}

// Find the minimum element in a non-empty list.
func Minimum[T Basics.Ordered](list []T) Maybe[T] {
	if len(list) == 0 {
		return Nothing[T]()
	}
	return Just(Basics.Min(list[0], list[1:]...))
}

func MaximumIndex[T Basics.Ordered](list []T) Maybe[int] {
	if len(list) == 0 {
		return Nothing[int]()
	}
	max_index := 0
	max_val := list[max_index]
	for i, val := range list[1:] {
		if val > max_val {
			max_val = val
			max_index = i
		}
	}
	return Just(max_index)
}

func MinimumIndex[T Basics.Ordered](list []T) Maybe[int] {
	if len(list) == 0 {
		return Nothing[int]()
	}
	min_index := 0
	min_val := list[min_index]
	for i, val := range list {
		if val < min_val {
			min_val = val
			min_index = i
		}
	}
	return Just(min_index)
}

// Get the sum of the list elements.
func Sum[T Basics.Number](list []T) Maybe[T] {
	if len(list) == 0 {
		return Nothing[T]()
	}
	return Just(Foldl(Basics.Add[T], T(0), list))
}

// Get the product of the list elements.
func Product[T Basics.Number](list []T) Maybe[T] {
	if len(list) == 0 {
		return Nothing[T]()
	}
	return Just(Foldl(Basics.Mul[T], T(0), list))
}

// COMBINE

// Put two lists together.
// This functions is IMMUTABLE and produces a completely new list!
func Append[T any](listA []T, listB []T) []T {
	return append(clone(listA), listB...)
}

// Concatenate a bunch of lists into a single list:
// This functions is IMMUTABLE and produces a completely new list!
func Concat[T any](list [][]T) []T {
	return Foldr(Append[T], []T{}, list)
}

// Map a given function onto a list and flatten the resulting lists.
// This functions is IMMUTABLE and produces a completely new list!
func ConcatMap[T, U any](mapfn func(value T) []U, list []T) []U {
	return Concat(Map(mapfn, list))
}

// Places the given value between all members of the given list.
// This functions is IMMUTABLE and produces a completely new list!
func Intersperse[T any](value T, list []T) []T {
	length := len(list)
	if length < 2 {
		return list
	}
	new_list := make([]T, 0, 2*length-1)
	for i, v := range list {
		new_list = append(new_list, v)
		if i != length-1 {
			new_list = append(new_list, value)
		}
	}
	return new_list
}

// Combine two lists, combining them with the given function. If one list is longer, the extra elements are dropped.
// This functions is IMMUTABLE and produces a completely new list!
func Map2[a, b, result any](mapfn func(a a, b b) result, listA []a, listB []b) []result {
	length := Basics.Min(len(listA), len(listB))
	new_list := make([]result, 0, length)
	for i := 0; i < length; i++ {
		new_list = append(new_list, mapfn(listA[i], listB[i]))
	}
	return new_list
}

// Combine three lists, combining them with the given function. If one list is longer, the extra elements are dropped.
// This functions is IMMUTABLE and produces a completely new list!
func Map3[a, b, c, result any](mapfn func(a a, b b, c c) result, lista []a, listb []b, listc []c) []result {
	length := Basics.Min(len(lista), len(listb), len(listc))
	new_list := make([]result, 0, length)
	for i := 0; i < length; i++ {
		new_list = append(new_list, mapfn(lista[i], listb[i], listc[i]))
	}
	return new_list
}

// Combine four lists, combining them with the given function. If one list is longer, the extra elements are dropped.
// This functions is IMMUTABLE and produces a completely new list!
func Map4[a, b, c, d, result any](mapfn func(a a, b b, c c, d d) result, lista []a, listb []b, listc []c, listd []d) []result {
	length := Basics.Min(len(lista), len(listb), len(listc), len(listd))
	new_list := make([]result, 0, length)
	for i := 0; i < length; i++ {
		new_list = append(new_list, mapfn(lista[i], listb[i], listc[i], listd[i]))
	}
	return new_list
}

// Combine five lists, combining them with the given function. If one list is longer, the extra elements are dropped.
// This functions is IMMUTABLE and produces a completely new list!
func Map5[a, b, c, d, e, result any](mapfn func(a a, b b, c c, d d, e e) result, lista []a, listb []b, listc []c, listd []d, liste []e) []result {
	length := Basics.Min(len(lista), len(listb), len(listc), len(listd), len(liste))
	new_list := make([]result, 0, length)
	for i := 0; i < length; i++ {
		new_list = append(new_list, mapfn(lista[i], listb[i], listc[i], listd[i], liste[i]))
	}
	return new_list
}

// SORT

// Sort values from lowest to highest
// This functions is IMMUTABLE and produces a completely new list!
func Sort[T Basics.Ordered](list []T) []T {
	new_list := clone(list)
	n := len(new_list)
	pdqsortOrdered(new_list, 0, n, bits.Len(uint(n)))
	return new_list
}

// Sort values by a derived property.
// This functions is IMMUTABLE and produces a completely new list!
func SortBy[T any, U Basics.Ordered](mapfn func(value T) U, list []T) []T {
	new_list := clone(list)
	n := len(new_list)
	pdqsortCmpFunc(new_list, 0, n, bits.Len(uint(n)), func(a, b T) Basics.Order {
		return Basics.Compare(mapfn(a), mapfn(b))
	})
	return new_list
}

// Sort values with a custom comparison function.
// This functions is IMMUTABLE and produces a completely new list!
func SortWith[T any](cmpfn func(a, b T) Basics.Order, list []T) []T {
	new_list := clone(list)
	n := len(new_list)
	pdqsortCmpFunc(new_list, 0, n, bits.Len(uint(n)), cmpfn)
	return new_list
}

// DECONSTRUCT

// Determine if a list is empty.
func IsEmpty[T any](list []T) bool {
	return len(list) == 0
}

// Extract the first element of a list.
func Head[T any](list []T) Maybe[T] {
	if len(list) == 0 {
		return Nothing[T]()
	}
	return Just(list[0])
}

// Extract the rest of the list.
func Tail[T any](list []T) Maybe[[]T] {
	if len(list) == 0 {
		return Nothing[[]T]()
	}
	return Just(list[1:])
}

// Take the first n members of a list.
func Take[N Basics.Int, T any](n N, list []T) []T {
	return list[0:n]
}

// Drop the first n members of a list.
func Drop[N Basics.Int, T any](n N, list []T) []T {
	return list[n:]
}

// Partition a list based on some test. The first list contains all values that satisfy the test, and the second list contains all the value that do not.
func Partition[T any](testfn func(value T) bool, list []T) ([]T, []T) {
	length := len(list)
	new_listA := make([]T, 0, length)
	new_listB := make([]T, 0, length)
	for _, v := range list {
		if testfn(v) {
			new_listA = append(new_listA, v)
		} else {
			new_listB = append(new_listB, v)
		}
	}
	return new_listA, new_listB
}

// Decompose a list of tuples into a tuple of lists.
func Unzip[T, U any](list []Tuple.Tuple[T, U]) Tuple.Tuple[[]T, []U] {
	list1 := make([]T, len(list))
	list2 := make([]U, len(list))
	for i, v := range list {
		list1[i] = v.Fst
		list2[i] = v.Snd
	}
	return Tuple.Pair(list1, list2)
}

// FROM ARRAY

// CREATE

// Initialize an array. initialize n f creates an array of length n with the element at index i initialized to the result of (f i).
func Initialize[T any](n int, mapfn func(index int) T) []T {
	return Map(mapfn, Range(0, n-1))
}

// Return Just the element at the index or Nothing if the index is out of range.
func Get[T any](index int, list []T) Maybe[T] {
	if index >= len(list) || index < 0 {
		return Nothing[T]()
	}
	return Just(list[index])
}

// Set the element at a particular index. Returns an updated array.
// If the index is out of range, the array is unaltered.
// This functions is IMMUTABLE and produces a completely new list!
func Set[T any](index int, value T, list []T) []T {
	if index < 0 || index >= len(list) {
		return list
	}
	new_list := clone(list)
	new_list[index] = value
	return new_list
}

// Push an element onto the end of an array.
// This functions is IMMUTABLE and produces a completely new list!
func Push[T any](value T, list []T) []T {
	return append(clone(list), value)
}

// Get a sub-section of an array: (slice start end array).
// The start is a zero-based index where we will start our slice.
// The end is a zero-based index that indicates the end of the slice.
// The slice extracts up to but not including end.
//
// Both the start and end indexes can be negative, indicating an offset from the end of the array.
//
// This makes it pretty easy to pop the last element off of an array: slice 0 -1 array
func Slice[T any](start int, stop int, list []T) []T {
	length := len(list)
	if start < 0 {
		start = Basics.Max(length+start, 0)
	}
	if stop < 0 {
		stop = Basics.Min(length+stop, length)
	}
	if stop <= 0 || start >= stop {
		return []T{}
	}
	return list[start:stop]
}

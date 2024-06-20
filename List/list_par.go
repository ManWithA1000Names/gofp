package List

import (
	"sync"

	"github.com/manwitha1000names/gofp/v3/Basics"
	. "github.com/manwitha1000names/gofp/v3/MaybeResult"
)

// TRANSFORM

// Apply a function to every element of a list.
// This functions is IMMUTABLE and produces a completely new list!
// Ordering IS preserved!
func Map_par[T, U any](mapfn func(value T) U, list []T) []U {
	new_list := make([]U, len(list))
	var wg sync.WaitGroup
	wg.Add(len(list))
	for i, w := range list {
		go func(i int, w T) {
			defer wg.Done()
			new_list[i] = mapfn(w)
		}(i, w)
	}
	wg.Wait()
	return new_list
}

// Same as map but the function is also applied to the index of each element (starting at zero).
// This functions is IMMUTABLE and produces a completely new list!
// Ordering IS preserved!
func IndexedMap_par[T, U any](mapfn func(index int, value T) U, list []T) []U {
	new_list := make([]U, len(list))
	var wg sync.WaitGroup
	wg.Add(len(list))
	for i, w := range list {
		go func(i int, w T) {
			defer wg.Done()
			new_list[i] = mapfn(i, w)
		}(i, w)
	}
	wg.Wait()
	return new_list
}

// Keep elements that satisfy the test.
// This functions is IMMUTABLE and produces a completely new list!
// Ordering IS NOT preserved!
func Filter_par[T any](testfn func(value T) bool, list []T) []T {
	length := len(list)
	new_list := make([]T, 0, length)
	ch := make(chan T, length)

	var wg sync.WaitGroup
	wg.Add(length)

	for _, v := range list {
		go func(v T) {
			defer wg.Done()
			if testfn(v) {
				ch <- v
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		new_list = append(new_list, value)
	}

	return new_list
}

// Filter out certain values.
// This functions is IMMUTABLE and produces a completely new list!
// Ordering IS NOT preserved!
func FilterMap_par[T, U any](testmapfn func(value T) Maybe[U], list []T) []U {
	length := len(list)
	new_list := make([]U, 0, length)
	ch := make(chan U, length)

	var wg sync.WaitGroup
	wg.Add(length)

	for _, v := range list {
		go func(v T) {
			defer wg.Done()
			testmapfn(v).Then(func(value U) { ch <- value })
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		new_list = append(new_list, value)
	}

	return new_list
}

// COMBINE

// Map a given function onto a list and flatten the resulting lists.
// This functions is IMMUTABLE and produces a completely new list!
// Ordering IS preserved!
func ConcatMap_par[T, U any](mapfn func(value T) []U, list []T) []U {
	return Concat(Map_par(mapfn, list))
}

// Combine two lists, combining them with the given function. If one list is longer, the extra elements are dropped.
// This functions is IMMUTABLE and produces a completely new list!
// Ordering IS preserved!
func Map2_par[a, b, result any](mapfn func(a a, b b) result, listA []a, listB []b) []result {
	length := Basics.Min(len(listA), len(listB))
	new_list := make([]result, length)
	var wg sync.WaitGroup
	wg.Add(length)
	for i := 0; i < length; i++ {
		go func(i int) {
			defer wg.Done()
			new_list[i] = mapfn(listA[i], listB[i])
		}(i)
	}
	wg.Wait()
	return new_list
}

// Combine three lists, combining them with the given function. If one list is longer, the extra elements are dropped.
// This functions is IMMUTABLE and produces a completely new list!
// Ordering IS preserved!
func Map3_par[a, b, c, result any](mapfn func(a a, b b, c c) result, lista []a, listb []b, listc []c) []result {
	length := Basics.Min(len(lista), len(listb), len(listc))
	new_list := make([]result, length)
	var wg sync.WaitGroup
	wg.Add(length)
	for i := 0; i < length; i++ {
		go func(i int) {
			defer wg.Done()
			new_list[i] = mapfn(lista[i], listb[i], listc[i])
		}(i)
	}
	wg.Wait()
	return new_list
}

// Combine four lists, combining them with the given function. If one list is longer, the extra elements are dropped.
// This functions is IMMUTABLE and produces a completely new list!
// Ordering IS preserved!
func Map4_par[a, b, c, d, result any](mapfn func(a a, b b, c c, d d) result, lista []a, listb []b, listc []c, listd []d) []result {
	length := Basics.Min(len(lista), len(listb), len(listc), len(listd))
	new_list := make([]result, length)
	var wg sync.WaitGroup
	wg.Add(length)
	for i := 0; i < length; i++ {
		go func(i int) {
			new_list[i] = mapfn(lista[i], listb[i], listc[i], listd[i])
		}(i)
	}
	wg.Wait()
	return new_list
}

// Combine five lists, combining them with the given function. If one list is longer, the extra elements are dropped.
// This functions is IMMUTABLE and produces a completely new list!
func Map5_par[a, b, c, d, e, result any](mapfn func(a a, b b, c c, d d, e e) result, lista []a, listb []b, listc []c, listd []d, liste []e) []result {
	length := Basics.Min(len(lista), len(listb), len(listc), len(listd), len(liste))
	new_list := make([]result, 0, length)
	var wg sync.WaitGroup
	wg.Add(length)
	for i := 0; i < length; i++ {
		go func(i int) {
			defer wg.Done()
			new_list[i] = mapfn(lista[i], listb[i], listc[i], listd[i], liste[i])
		}(i)
	}
	wg.Wait()
	return new_list
}

// Partition a list based on some test. The first list contains all values that satisfy the test, and the second list contains all the value that do not.
func Partition_par[T any](testfn func(value T) bool, list []T) ([]T, []T) {
	length := len(list)
	new_listA := make([]T, 0, length)
	new_listB := make([]T, 0, length)
	var wg sync.WaitGroup
	wg.Add(len(list))

	chA := make(chan T, length)
	chB := make(chan T, length)
	chDone := make(chan bool)

	for _, v := range list {
		go func(v T) {
			if testfn(v) {
				chA <- v
			} else {
				chB <- v
			}
		}(v)
	}

	go func() {
		wg.Wait()
		chDone <- true
	}()

	for {
		select {
		case v := <-chA:
			new_listA = append(new_listA, v)
		case v := <-chB:
			new_listB = append(new_listB, v)
		case <-chDone:
			return new_listA, new_listB
		}
	}
}

// FROM ARRAY

// Initialize an array. initialize n f creates an array of length n with the element at index i initialized to the result of (f i).
func Initialize_par[T any](n int, mapfn func(index int) T) []T {
	return Map_par(mapfn, Range(0, n-1))
}

package Set

import (
	"sync"

	"github.com/manwitha1000names/gofp/Dict"
	"github.com/manwitha1000names/gofp/List"
)

// TRANSFORM

// Map a function onto a set, creating a new set with no duplicates.
// This functions is IMMUTABLE and produces a completely new set!
func Map_par[T comparable, U comparable](mapfn func(value T) U, s Set[T]) Set[U] {
	length := len(s.m)
	new_map := make(map[U]struct{}, length)
	ch := make(chan U, length)
	var wg sync.WaitGroup
	wg.Add(length)
	for key := range s.m {
		go func(key T) {
			defer wg.Done()
			ch <- mapfn(key)
		}(key)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for key := range ch {
		new_map[key] = struct{}{}
	}

	return Set[U]{new_map}
}

// Only keep elements that pass the given test.
// This functions is IMMUTABLE and produces a completely new set!
// There is absolutely no guarantee on any order.
func Filter_par[T comparable](testfn func(value T) bool, s Set[T]) Set[T] {
	return Set[T]{Dict.Filter_par(func(value T, _ struct{}) bool {
		return testfn(value)
	}, s.m)}
}

// Create two new sets.
// The first contains all the elements that passed the given test,
// and the second contains all the elements that did not.
func Partition_par[T comparable](partfn func(value T) bool, s Set[T]) (Set[T], Set[T]) {
	lista, listb := List.Partition_par(partfn, ToList(s))
	return FromList(lista), FromList(listb)
}

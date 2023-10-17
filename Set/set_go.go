package Set

import "sync"

// Perform an action for each element of a set.
// A.K.A Side-effect heaven.
func ForEach[T comparable](fn func(value T), s Set[T]) {
	for value := range s.m {
		fn(value)
	}
}

// Perform an action for each element of a set in parallel.
// A.K.A Parallel side-effect heaven.
func ForEach_par[T comparable](fn func(value T), s Set[T]) {
	var wg sync.WaitGroup
	wg.Add(len(s.m))
	for value := range s.m {
		go func(value T) {
			defer wg.Done()
			fn(value)
		}(value)
	}
	wg.Wait()
}

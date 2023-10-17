package List

import "sync"

// Perform an action for each element of a list.
// A.K.A Side-effect heaven.
func ForEach[T any](fn func(value *T), list []T) {
	for i := 0; i < len(list); i++ {
		fn(&list[i])
	}
}

// Perform an action for each element of a list in parallel.
// A.K.A Parallel side-effect heaven.
func ForEach_par[T any](fn func(value *T), list []T) {
	var wg sync.WaitGroup
	wg.Add(len(list))
	for i := 0; i < len(list); i++ {
		go func(i int) {
			defer wg.Done()
			fn(&list[i])
		}(i)
	}
	wg.Wait()
}

// Like the `Any` function but the first value found is returned
// The pointer returned, points directly into the given array!
func Find[T any](testfn func(value T) bool, list []T) *T {
	for i := 0; i < len(list); i++ {
		if testfn(list[i]) {
			return &list[i]
		}
	}
	return nil
}

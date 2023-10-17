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

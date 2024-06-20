package Dict

import (
	"sync"

	. "github.com/manwitha1000names/gofp/v3/MaybeResult"
)

// Perform an action for each element of a dict.
// A.K.A Side-effect heaven.
func ForEach[Key comparable, Value any](fn func(key Key, value Value), m map[Key]Value) {
	for key, value := range m {
		fn(key, value)
	}
}

// Perform an action for each element of a dict in parallel.
// A.K.A Parallel side-effect heaven.
func ForEach_par[Key comparable, Value any](fn func(key Key, value Value), m map[Key]Value) {
	var wg sync.WaitGroup
	wg.Add(len(m))
	for key, value := range m {
		go func(key Key, value Value) {
			defer wg.Done()
			fn(key, value)
		}(key, value)
	}
	wg.Wait()
}

// Find the first value found that passes the testfn and return it.
func Find[Key comparable, Value any](testfn func(key Key, value Value) bool, dict map[Key]Value) Maybe[Value] {
	for key, value := range dict {
		if testfn(key, value) {
			return Just(value)
		}
	}
	return Nothing[Value]()
}

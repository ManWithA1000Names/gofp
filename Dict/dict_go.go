package Dict

import "sync"

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
// The pointer returned, points to a copy of the value.
func Find[Key comparable, Value any](testfn func(key Key, value Value) bool, dict map[Key]Value) *Value {
	for key, value := range dict {
		if testfn(key, value) {
			return &value
		}
	}
	return nil
}

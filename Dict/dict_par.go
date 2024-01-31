package Dict

import "sync"

// TRANSFORM

// Apply a function to all values in a dictionary.
// This functions is IMMUTABLE and produces a completely new map!
func Map_par[Key comparable, Value1 any, Value2 any](mapfn func(key Key, value Value1) Value2, m map[Key]Value1) map[Key]Value2 {
	new_map := make(map[Key]Value2, len(m))
	var wg sync.WaitGroup
	wg.Add(len(m))
	for _, key := range Keys(m) {
		go func(key Key) {
			defer wg.Done()
			new_map[key] = mapfn(key, m[key])
		}(key)
	}
	wg.Wait()
	return new_map
}

// Keep only the key-value pairs that pass the given test.
// This functions is IMMUTABLE and produces a completely new map!
func Filter_par[Key comparable, Value any](testfn func(key Key, value Value) bool, m map[Key]Value) map[Key]Value {
	new_map := make(map[Key]Value, len(m))
	var wg sync.WaitGroup
	wg.Add(len(m))
	for key, value := range m {
		go func(key Key, value Value) {
			defer wg.Done()
			if testfn(key, value) {
				new_map[key] = value
			}
		}(key, value)
	}
	wg.Wait()
	return new_map
}

// Partition a dictionary according to some test. The first dictionary contains all key-value pairs which passed the test, and the second contains the pairs that did not.
func Partition_par[Key comparable, Value any](partfn func(key Key, value Value) bool, m map[Key]Value) (map[Key]Value, map[Key]Value) {
  length := len(m)
	map1 := make(map[Key]Value, length)
	map2 := make(map[Key]Value, length)
	var wg sync.WaitGroup
	wg.Add(length)
	for key, value := range m {
		go func(key Key, value Value) {
			defer wg.Done()
			if partfn(key, value) {
				map1[key] = value
			} else {
				map2[key] = value
			}
		}(key, value)
	}
	wg.Wait()
	return map1, map2
}

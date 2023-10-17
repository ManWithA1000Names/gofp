package String

import "sync"

// Perform an action for each character of a string.
// A.K.A Side-effect heaven.
func ForEach(fn func(c rune), s string) {
	for _, value := range s {
		fn(value)
	}
}

// Perform an action for each character of a sstring in parallel.
// A.K.A Parallel side-effect heaven.
func ForEach_par(fn func(c rune), s string) {
	var wg sync.WaitGroup
	runes := []rune(s)
	wg.Add(len(runes))
	for _, c := range runes {
		go func(c rune) {
			defer wg.Done()
			fn(c)
		}(c)
	}
	wg.Wait()
}

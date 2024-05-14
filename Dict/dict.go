package Dict

import (
	"github.com/manwitha1000names/gofp/v2/List"
	. "github.com/manwitha1000names/gofp/v2/MaybeResult"
	"github.com/manwitha1000names/gofp/v2/Tuple"
)

func clone[Key comparable, Value any](m map[Key]Value) map[Key]Value {
	new_map := make(map[Key]Value, len(m))
	for key, value := range m {
		new_map[key] = value
	}
	return new_map
}

// BUILD

// Create an empty dictionary.
func Empty[Key comparable, Value any]() map[Key]Value {
	return make(map[Key]Value)
}

// Create a dictionary with one key-value pair.
func Singleton[Key comparable, Value any](key Key, value Value) map[Key]Value {
	m := make(map[Key]Value)
	m[key] = value
	return m
}

// Insert a key-value pair into a dictionary. Replaces value when there is a collision.
// This functions is IMMUTABLE and produces a completely new map!
func Insert[Key comparable, Value any](key Key, v Value, m map[Key]Value) map[Key]Value {
	new_map := clone(m)
	new_map[key] = v
	return new_map
}

// Update the value of a dictionary for a specific key with a given function.
// This functions is IMMUTABLE and produces a completely new map!
func Update[Key comparable, Value any](key Key, upfn func(value Value) Value, m map[Key]Value) map[Key]Value {
	if value, ok := m[key]; ok {
		new_map := clone(m)
		new_map[key] = upfn(value)
		return new_map
	}
	return m
}

// Remove a key-value pair from a dictionary. If the key is not found, no changes are made.
// This functions is IMMUTABLE and produces a completely new map!
func Remove[Key comparable, Value any](key Key, m map[Key]Value) map[Key]Value {
	if _, ok := m[key]; ok {
		new_map := clone(m)
		delete(new_map, key)
		return new_map
	}
	return m
}

// QUERY

// Determine if a dictionary is empty.
func IsEmpty[Key comparable, Value any](m map[Key]Value) bool {
	return len(m) == 0
}

// Determine if a key is in a dictionary.
func Member[Key comparable, Value any](key Key, m map[Key]Value) bool {
	_, ok := m[key]
	return ok
}

// Get the value associated with a key. If the key is not found, return Nothing.
// This is useful when you are not sure if a key will be in the dictionary.
func Get[Key comparable, Value any](key Key, m map[Key]Value) Maybe[Value] {
	value, ok := m[key]
	return TupleToMaybe(value, ok)
}

// Determine the number of key-value pairs in the dictionary.
func Size[Key comparable, Value any](m map[Key]Value) int {
	return len(m)
}

// LISTS

// Get all of the keys in a dictionary, NOT IN ANY PARTICULAR ORDER.
func Keys[Key comparable, Value any](m map[Key]Value) []Key {
	return List.Map(Tuple.First[Key, Value], ToList(m))
}

// Get all of the values in a dictionary, NOT IN ANY PARTICULAR ORDER.
func Values[Key comparable, Value any](m map[Key]Value) []Value {
	return List.Map(Tuple.Second[Key, Value], ToList(m))
}

// Convert a dictionary into an association list of key-value pairs, NOT IN ANY PARTICULAR ORDER.
func ToList[Key comparable, Value any](m map[Key]Value) []Tuple.Tuple[Key, Value] {
	list := make([]Tuple.Tuple[Key, Value], 0, len(m))
	for key, value := range m {
		list = append(list, Tuple.Pair(key, value))
	}
	return list
}

// Convert an association list into a dictionary.
func FromList[Key comparable, Value any](list []Tuple.Tuple[Key, Value]) map[Key]Value {
	m := make(map[Key]Value, len(list))
	for _, t := range list {
		m[Tuple.First(t)] = Tuple.Second(t)
	}
	return m
}

// TRANSFORM

// Apply a function to all values in a dictionary.
// This functions is IMMUTABLE and produces a completely new map!
func Map[Key comparable, Value1 any, Value2 any](mapfn func(key Key, value Value1) Value2, m map[Key]Value1) map[Key]Value2 {
	new_map := make(map[Key]Value2, len(m))
	for _, key := range Keys(m) {
		new_map[key] = mapfn(key, m[key])
	}
	return new_map
}

// Fold over the key-value pairs in a dictionary, from left to right on the resulting list from calling Dict.ToList.
func Foldl[Key comparable, Value any, Acc any](reducer func(key Key, value Value, acc Acc) Acc, init Acc, m map[Key]Value) Acc {
	return List.Foldl(func(t Tuple.Tuple[Key, Value], acc Acc) Acc {
		return reducer(t.Fst, t.Snd, acc)
	}, init, ToList(m))
}

// Fold over the key-value pairs in a dictionary, from right to left on the resulting list from calling Dict.ToList..
func Foldr[Key comparable, Value any, Acc any](recuder func(key Key, value Value, acc Acc) Acc, init Acc, m map[Key]Value) Acc {
	return List.Foldr(func(t Tuple.Tuple[Key, Value], acc Acc) Acc {
		return recuder(t.Fst, t.Snd, acc)
	}, init, ToList(m))
}

// Keep only the key-value pairs that pass the given test.
// This functions is IMMUTABLE and produces a completely new map!
func Filter[Key comparable, Value any](testfn func(key Key, value Value) bool, m map[Key]Value) map[Key]Value {
	new_map := make(map[Key]Value, len(m))
	for key, value := range m {
		if testfn(key, value) {
			new_map[key] = value
		}
	}
	return new_map
}

// Partition a dictionary according to some test. The first dictionary contains all key-value pairs which passed the test, and the second contains the pairs that did not.
func Partition[Key comparable, Value any](partfn func(key Key, value Value) bool, m map[Key]Value) (map[Key]Value, map[Key]Value) {
	map1 := make(map[Key]Value, len(m))
	map2 := make(map[Key]Value, len(m))
	for key, value := range m {
		if partfn(key, value) {
			map1[key] = value
		} else {
			map2[key] = value
		}
	}
	return map1, map2
}

// COMBINE

// Combine two dictionaries. If there is a collision, preference is given to the first dictionary.
// This functions is IMMUTABLE and produces a completely new map!
func Union[Key comparable, Value any](m map[Key]Value, m1 map[Key]Value) map[Key]Value {
	m2 := clone(m1)
	for key, value := range m {
		m2[key] = value
	}
	return m2
}

// Keep a key-value pair when its key appears in the second dictionary. Preference is given to values in the first dictionary.
// This functions is IMMUTABLE and produces a completely new map!
func Intersect[Key comparable, Value any](m map[Key]Value, m1 map[Key]Value) map[Key]Value {
	return Filter(func(key Key, _ Value) bool {
		return Member(key, m1)
	}, m)
}

// Keep a key-value pair when its key does not appear in the second dictionary.
// This functions is IMMUTABLE and produces a completely new map!
func Diff[Key comparable, Value any](m map[Key]Value, m1 map[Key]Value) map[Key]Value {
	return Filter(func(key Key, _ Value) bool {
		return !Member(key, m1)
	}, m)
}

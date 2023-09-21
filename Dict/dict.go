package Dict

import (
	"github.com/manwitha1000names/gofp/List"
	"github.com/manwitha1000names/gofp/Maybe"
	"github.com/manwitha1000names/gofp/Tuple"
)

func clone[Key comparable, Value any](m map[Key]Value) map[Key]Value {
	new_map := make(map[Key]Value, len(m))
	for key, value := range m {
		new_map[key] = value
	}
	return new_map
}

func Empty[Key comparable, Value any]() map[Key]Value {
	return make(map[Key]Value)
}

func Singleton[Key comparable, Value any](key Key, value Value) map[Key]Value {
	m := make(map[Key]Value)
	m[key] = value
	return m
}

func Insert[Key comparable, Value any](key Key, v Value, m map[Key]Value) map[Key]Value {
	new_map := clone(m)
	new_map[key] = v
	return new_map
}

func Update[Key comparable, Value any](key Key, upfn func(value Value) Value, m map[Key]Value) map[Key]Value {
	if value, ok := m[key]; ok {
		new_map := clone(m)
		new_map[key] = upfn(value)
		return new_map
	}
	return m
}

func Remove[Key comparable, Value any](key Key, m map[Key]Value) map[Key]Value {
	if _, ok := m[key]; ok {
		new_map := clone(m)
		delete(new_map, key)
		return new_map
	}
	return m
}

func IsEmpty[Key comparable, Value any](m map[Key]Value) bool {
	return len(m) == 0
}

func Member[Key comparable, Value any](key Key, m map[Key]Value) bool {
	_, ok := m[key]
	return ok
}

func Get[Key comparable, Value any](key Key, m map[Key]Value) Maybe.Maybe[Value] {
	if value, ok := m[key]; ok {
		return Maybe.Just(value)
	}
	return Maybe.Nothing[Value]()
}

func Size[Key comparable, Value any](m map[Key]Value) int {
	return len(m)
}

func Keys[Key comparable, Value any](m map[Key]Value) []Key {
	return List.Map(Tuple.First[Key, Value], ToList(m))
}

func Values[Key comparable, Value any](m map[Key]Value) []Value {
	return List.Map(Tuple.Second[Key, Value], ToList(m))
}

func ToList[Key comparable, Value any](m map[Key]Value) []Tuple.Tuple[Key, Value] {
	list := make([]Tuple.Tuple[Key, Value], 0, len(m))
	for key, value := range m {
		list = append(list, Tuple.Pair(key, value))
	}
	return list
	// return List.SortBy(Tuple.First[Key, Value], list) // AAAAH.
}

func FromList[Key comparable, Value any](list []Tuple.Tuple[Key, Value]) map[Key]Value {
	m := make(map[Key]Value, len(list))
	for _, t := range list {
		m[Tuple.First(t)] = Tuple.Second(t)
	}
	return m
}

func Map[Key comparable, Value1 any, Value2 any](mapfn func(value Value1) Value2, m map[Key]Value1) map[Key]Value2 {
	new_map := make(map[Key]Value2, len(m))
	for _, key := range Keys(m) {
		new_map[key] = mapfn(m[key])
	}
	return new_map
}

func Foldl[Key comparable, Value any, Acc any](reducer func(key Key, value Value, acc Acc) Acc, init Acc, m map[Key]Value) Acc {
	return List.Foldl(func(t Tuple.Tuple[Key, Value], acc Acc) Acc {
		return reducer(t.Fst, t.Snd, acc)
	}, init, ToList(m))
}

func Foldr[Key comparable, Value any, Acc any](recuder func(key Key, value Value, acc Acc) Acc, init Acc, m map[Key]Value) Acc {
	return List.Foldr(func(t Tuple.Tuple[Key, Value], acc Acc) Acc {
		return recuder(t.Fst, t.Snd, acc)
	}, init, ToList(m))
}

func Filter[Key comparable, Value any](testfn func(key Key, value Value) bool, m map[Key]Value) map[Key]Value {
	new_map := make(map[Key]Value, len(m))
	for key, value := range m {
		if testfn(key, value) {
			new_map[key] = value
		}
	}
	return new_map
}

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

func Union[Key comparable, Value any](m map[Key]Value, m1 map[Key]Value) map[Key]Value {
	m2 := clone(m1)
	for key, value := range m {
		m2[key] = value
	}
	return m2
}

func Intersect[Key comparable, Value any](m map[Key]Value, m1 map[Key]Value) map[Key]Value {
	return Filter(func(key Key, _ Value) bool {
		return Member(key, m1)
	}, m)
}

func Diff[Key comparable, Value any](m map[Key]Value, m1 map[Key]Value) map[Key]Value {
	return Filter(func(key Key, _ Value) bool {
		return !Member(key, m1)
	}, m)
}

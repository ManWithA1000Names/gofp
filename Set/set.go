package Set

import (
	"github.com/manwitha1000names/yggdrasil/fp/Dict"
	"github.com/manwitha1000names/yggdrasil/fp/List"
	"github.com/manwitha1000names/yggdrasil/fp/Tuple"
)

type Set[T comparable] struct {
	m map[T]int
}

// BUILD

func Empty[T comparable]() Set[T] {
	return Set[T]{make(map[T]int)}
}

func Singleton[T comparable](value T) Set[T] {
	new_map := make(map[T]int)
	new_map[value] = 0
	return Set[T]{new_map}
}

func Insert[T comparable](value T, s Set[T]) Set[T] {
	return Set[T]{Dict.Insert(value, 0, s.m)}
}

func Remove[T comparable](value T, s Set[T]) Set[T] {
	return Set[T]{Dict.Remove(value, s.m)}
}

// QUERY

func IsEmpty[T comparable](s Set[T]) bool {
	return Dict.IsEmpty(s.m)
}

func Member[T comparable](value T, s Set[T]) bool {
	return Dict.Member(value, s.m)
}

func Size[T comparable](s Set[T]) int {
	return len(s.m)
}

// COMBINE

func Union[T comparable](s Set[T], s1 Set[T]) Set[T] {
	return Set[T]{Dict.Union(s.m, s1.m)}
}

func Intersect[T comparable](s Set[T], s1 Set[T]) Set[T] {
	return Set[T]{Dict.Intersect(s.m, s1.m)}
}

func Diff[T comparable](s Set[T], s1 Set[T]) Set[T] {
	return Set[T]{Dict.Diff(s.m, s1.m)}
}

// LISTS

func ToList[T comparable](s Set[T]) []T {
	return Dict.Keys(s.m)
}

func FromList[T comparable](list []T) Set[T] {
	return Set[T]{Dict.FromList(List.Map(func(value T) Tuple.Tuple[T, int] {
		return Tuple.Pair(value, 0)
	}, list))}
}

// TRANSFORM

func Map[T comparable, U comparable](mapfn func(value T) U, s Set[T]) Set[U] {
	new_map := make(map[U]int, len(s.m))
	for key := range s.m {
		new_map[mapfn(key)] = 0
	}
	return Set[U]{new_map}
}

func Foldl[T comparable, Acc any](reducer func(value T, acc Acc) Acc, init Acc, s Set[T]) Acc {
	return List.Foldl(reducer, init, ToList(s))
}

func Foldr[T comparable, Acc any](reducer func(value T, acc Acc) Acc, init Acc, s Set[T]) Acc {
	return List.Foldr(reducer, init, ToList(s))
}

func Filter[T comparable](testfn func(value T) bool, s Set[T]) Set[T] {
	return FromList(List.Filter(testfn, ToList(s)))
}

func Partition[T comparable](partfn func(value T) bool, s Set[T]) (Set[T], Set[T]) {
	lista, listb := List.Partition(partfn, ToList(s))
	return FromList(lista), FromList(listb)
}

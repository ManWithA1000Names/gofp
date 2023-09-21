package List

// import "github.com/manwitha1000names/gofp/Basics"
import (
	"math/bits"

	"github.com/manwitha1000names/gofp/Basics"
	"github.com/manwitha1000names/gofp/Maybe"
)

func clone[T any](value []T) []T {
	new_list := make([]T, len(value))
	copy(new_list, value)
	return new_list
}

// CREATE

func Singleton[T any](value T) []T {
	return []T{value}
}

func Repeat[T any](amount int, value T) []T {
	list := make([]T, 0, amount)
	for i := 0; i < amount; i++ {
		list = append(list, value)
	}
	return list
}

// inclusive
func Range(start, end int) []int {
	if start > end {
		return []int{}
	}
	length := end - start + 1
	list := make([]int, 0, length)
	for i := 0; i < length; i++ {
		list = append(list, i+start)
	}
	return list
}

func Cons[T any](value T, list []T) []T {
	return append(Singleton(value), list...)
}

// TRANSFORM

func Map[T, U any](f func(value T) U, list []T) []U {
	new_list := make([]U, 0, len(list))
	for _, w := range list {
		new_list = append(new_list, f(w))
	}
	return new_list
}

func IndexedMap[T, U any](f func(index int, value T) U, list []T) []U {
	new_list := make([]U, 0, len(list))
	for i, w := range list {
		new_list = append(new_list, f(i, w))
	}
	return new_list
}

func Foldl[T, U any](f func(value T, accumulator U) U, acc U, list []T) U {
	for _, w := range list {
		acc = f(w, acc)
	}
	return acc
}

func Foldr[T, U any](f func(value T, accumulator U) U, acc U, list []T) U {
	for i := len(list) - 1; i >= 0; i-- {
		acc = f(list[i], acc)
	}
	return acc
}

func Filter[T any](test func(value T) bool, list []T) []T {
	new_list := make([]T, 0, len(list))
	for _, v := range list {
		if test(v) {
			new_list = append(new_list, v)
		}
	}
	return new_list
}

func FilterMap[T, U any](test func(value T) Maybe.Maybe[U], list []T) []U {
	new_list := make([]U, 0, len(list))
	for _, v := range list {
		res := test(v)
		if res.IsJust() {
			new_list = append(new_list, res.Unwrap())
		}
	}
	return new_list
}

// UTILITIES

func Length[T any](list []T) int {
	return len(list)
}

func Reverse[T any](list []T) []T {
	return Foldl(Cons[T], []T{}, list)
}

func Member[T comparable](value T, list []T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func All[T any](test func(value T) bool, list []T) bool {
	for _, v := range list {
		if !test(v) {
			return false
		}
	}
	return true
}

func Any[T any](test func(value T) bool, list []T) bool {
	for _, v := range list {
		if test(v) {
			return false
		}
	}
	return false
}

func Maximum[T Basics.Ordered](list []T) Maybe.Maybe[T] {
	if len(list) == 0 {
		return Maybe.Nothing[T]()
	}
	return Maybe.Just(Basics.Max(list[0], list[1:]...))
}

func Minimum[T Basics.Ordered](list []T) Maybe.Maybe[T] {
	if len(list) == 0 {
		return Maybe.Nothing[T]()
	}
	return Maybe.Just(Basics.Min(list[0], list[1:]...))
}

func Sum[T Basics.Number](list []T) Maybe.Maybe[T] {
	if len(list) == 0 {
		return Maybe.Nothing[T]()
	}
	return Maybe.Just(Foldl(Basics.Add[T], T(0), list))
}

func Product[T Basics.Number](list []T) Maybe.Maybe[T] {
	if len(list) == 0 {
		return Maybe.Nothing[T]()
	}
	return Maybe.Just(Foldl(Basics.Mul[T], T(0), list))
}

func Append[T any](listA []T, listB []T) []T {
	return append(listA, listB...)
}

func Concat[T any](list [][]T) []T {
	return Foldr(Append[T], []T{}, list)
}

func ConcatMap[T, U any](mapfn func(value T) []U, list []T) []U {
	return Concat(Map(mapfn, list))
}

func Intersperse[T any](value T, list []T) []T {
	if len(list) < 2 {
		return list
	}
	length := len(list)
	new_list := make([]T, 0, length)
	for i, v := range list {
		new_list = append(new_list, v)
		if i != length-1 {
			new_list = append(new_list, value)
		}
	}
	return new_list
}

func Map2[a, b, result any](mapfn func(a a, b b) result, listA []a, listB []b) []result {
	length := Basics.Min(len(listA), len(listB))
	new_list := make([]result, 0, length)
	for i := 0; i < length; i++ {
		new_list = append(new_list, mapfn(listA[i], listB[i]))
	}
	return new_list
}

func Map3[a, b, c, result any](mapfn func(a a, b b, c c) result, lista []a, listb []b, listc []c) []result {
	length := Basics.Min(len(lista), len(listb), len(listc))
	new_list := make([]result, 0, length)
	for i := 0; i < length; i++ {
		new_list = append(new_list, mapfn(lista[i], listb[i], listc[i]))
	}
	return new_list
}

func Map4[a, b, c, d, result any](mapfn func(a a, b b, c c, d d) result, lista []a, listb []b, listc []c, listd []d) []result {
	length := Basics.Min(len(lista), len(listb), len(listc), len(listd))
	new_list := make([]result, 0, length)
	for i := 0; i < length; i++ {
		new_list = append(new_list, mapfn(lista[i], listb[i], listc[i], listd[i]))
	}
	return new_list
}

func Map5[a, b, c, d, e, result any](mapfn func(a a, b b, c c, d d, e e) result, lista []a, listb []b, listc []c, listd []d, liste []e) []result {
	length := Basics.Min(len(lista), len(listb), len(listc), len(listd), len(liste))
	new_list := make([]result, 0, length)
	for i := 0; i < length; i++ {
		new_list = append(new_list, mapfn(lista[i], listb[i], listc[i], listd[i], liste[i]))
	}
	return new_list
}

func Sort[T Basics.Ordered](list []T) []T {
	new_list := clone(list)
	n := len(new_list)
	pdqsortOrdered(new_list, 0, n, bits.Len(uint(n)))
	return new_list
}

func SortBy[T any, U Basics.Ordered](mapfn func(value T) U, list []T) []T {
	new_list := clone(list)
	n := len(new_list)
	pdqsortCmpFunc(new_list, 0, n, bits.Len(uint(n)), func(a, b T) Basics.Order {
		return Basics.Compare(mapfn(a), mapfn(b))
	})
	return new_list
}

func SortWith[T any](cmpfn func(a, b T) Basics.Order, list []T) []T {
	new_list := clone(list)
	n := len(new_list)
	pdqsortCmpFunc(new_list, 0, n, bits.Len(uint(n)), cmpfn)
	return new_list
}

func IsEmpty[T any](list []T) bool {
	return len(list) == 0
}

func Head[T any](list []T) Maybe.Maybe[T] {
	if len(list) == 0 {
		return Maybe.Nothing[T]()
	}
	return Maybe.Just(list[0])
}

func Tail[T any](list []T) Maybe.Maybe[[]T] {
	if len(list) == 0 {
		return Maybe.Nothing[[]T]()
	}
	return Maybe.Just(list[1:])
}

func Take[N Basics.Int, T any](n N, list []T) []T {
	return list[0:n]
}

func Drop[N Basics.Int, T any](n N, list []T) []T {
	return list[n:]
}

func Partition[T any](testfn func(value T) bool, list []T) ([]T, []T) {
	length := len(list)
	new_listA := make([]T, 0, length)
	new_listB := make([]T, 0, length)
	for _, v := range list {
		if testfn(v) {
			new_listA = append(new_listA, v)
		} else {
			new_listB = append(new_listB, v)
		}
	}
	return new_listA, new_listB
}

// FROM ARRAY

func Initialize[T any](n int, mapfn func(index int) T) []T {
	return Map(mapfn, Range(0, n-1))
}

func Get[T any](index int, list []T) Maybe.Maybe[T] {
	if index >= len(list) || index < 0 {
		return Maybe.Nothing[T]()
	}
	return Maybe.Just(list[index])
}

func Set[T any](index int, value T, list []T) []T {
	if index < 0 || index >= len(list) {
		return list
	}
	new_list := clone(list)
	new_list[index] = value
	return new_list
}

func Push[T any](value T, list []T) []T {
	new_list := clone(list)
	return append(new_list, value)
}

// supports negative indeces
func Slice[T any](start int, stop int, list []T) []T {
	if start < 0 {
		start = Basics.Max(len(list)+start, 0)
	}
	if stop < 0 {
		stop = Basics.Min(len(list)+stop, len(list))
	}
	if stop <= 0 || start >= stop {
		return []T{}
	}
	return list[start:stop]
}

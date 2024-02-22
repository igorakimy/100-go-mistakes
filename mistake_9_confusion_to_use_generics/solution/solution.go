package main

import (
	"fmt"
	"sort"
)

// merge функция обхединения двух каналов будет работать
// с любым типом канала.
func merge[T any](ch1, ch2 <-chan T) <-chan T {
	// ...
	return nil
}

// SliceFn использует параметр типа
type SliceFn[T any] struct {
	S []T
	// Сравниваются два элемента T
	Compare func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.S)
}

func (s SliceFn[T]) Less(i, j int) bool {
	return s.Compare(s.S[i], s.S[j])
}

func (s SliceFn[T]) Swap(i, j int) {
	s.S[i], s.S[j] = s.S[j], s.S[i]
}

func main() {
	s := SliceFn[int]{
		S: []int{3, 2, 1},
		Compare: func(a int, b int) bool {
			return a < b
		},
	}
	sort.Sort(s)
	fmt.Println(s.S)
	// [1 2 3]
}

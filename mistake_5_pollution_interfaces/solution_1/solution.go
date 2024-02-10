package main

import (
	"fmt"
	"sort"
)

// Пример хорошего интерфейса и его правильное использование

func isSorted(data sort.Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

func main() {
	sortable := isSorted(sort.IntSlice{1, 2, 3, 4})
	fmt.Println(sortable)
}

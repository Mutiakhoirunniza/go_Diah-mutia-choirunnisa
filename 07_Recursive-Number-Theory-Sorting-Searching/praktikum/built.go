package main

import (
	"fmt"
	"sort"
)

func builtin(elements []int, x int) int {
	// Menggunakan pencarian biner dalam urutan terurut
	return sort.Search(len(elements), func(i int) bool {
		return elements[i] >= x
	})
}

func main() {
	elements := []int{5, 2, 7, 1, 8, 3, 9, 4}
	searchValues := []int{12,7}

	for _, x := range searchValues {
		index := builtin(elements, x)
		if index < len(elements) && elements[index] == x {
			fmt.Printf("%d ditemukan pada indeks %d\n", x, index)
		} else {
			fmt.Printf("%d tidak ditemukan, akan disisipkan pada indeks %d\n", x, -index-1)
		}
	}
}

package main

import "fmt"

func linearSearch(elements []int, x int) int {
	for i := 0; i < len(elements); i++ {
		if elements[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	elements := []int{5, 2, 7, 1, 8, 3, 9, 4}
	searchValues := []int{12, 7}

	for _, x := range searchValues {
		indeks := linearSearch(elements, x)
		fmt.Printf("%d = %d\n", x, indeks)
	}
}

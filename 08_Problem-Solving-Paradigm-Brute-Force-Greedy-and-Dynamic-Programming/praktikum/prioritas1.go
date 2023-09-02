package main

import "fmt"

func generateIndices(n int) []int {
	indices := make([]int, n+1)

	for i := 0; i <= n; i++ {
		indices[i] = i
	}

	return indices
}

func main() {
	n1 := 2
	indices1 := generateIndices(n1)
	fmt.Println(indices1) // Output: [0 1 2]

	n2 := 3
	indices2 := generateIndices(n2)
	fmt.Println(indices2) // Output: [0 1 2 3]
}

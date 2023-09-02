package main

import (
	"fmt"
	"strconv"
)

// generateBinaryNumbers menghasilkan representasi biner dari bilangan bulat dari 0 hingga n.
func generateBinaryNumbers(n int) []string {
	ans := make([]string, n+1)

	for i := 0; i <= n; i++ {
		binaryString := strconv.FormatInt(int64(i), 2)
		ans[i] = binaryString
	}

	return ans
}

func main() {
	// Menggunakan generateBinaryNumbers untuk N1
	N1 := 2
	result1 := generateBinaryNumbers(N1)
	fmt.Printf("input:%d\n", N1)
	fmt.Printf("output %v\n", result1)

	// Menggunakan generateBinaryNumbers untuk N2
	N2 := 3
	result2 := generateBinaryNumbers(N2)
	fmt.Printf("input:%d\n", N2)
	fmt.Printf("output %v\n", result2)
}

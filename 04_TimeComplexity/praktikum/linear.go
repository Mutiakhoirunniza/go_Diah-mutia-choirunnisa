package main

import "fmt"

func linier2(n int, m int) int {
	result := 0
	for i := 0; i < n; i++ {
		result += 1
	}
	for j := 0; j < m; j++ {
		result += 1
	}
	return result
}

func main() {
	result := linier2(5, 3)
	fmt.Println(result)
}

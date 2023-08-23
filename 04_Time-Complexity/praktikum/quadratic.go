package main

import "fmt"

func quadratic(n int) int {
	var result int = 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			result += 1
		}
	}
	return result
}

func main() {
	result := quadratic(15)
	fmt.Println(result)
}

package main

import "fmt"

func dominant(n int) int {
	var result int = 0
	for i := 0; i < n; i++ {
		result += 1
	}
	result = result + 10
	return result
}

func mainn() {
	result := dominant(5)
	fmt.Println(result)
}

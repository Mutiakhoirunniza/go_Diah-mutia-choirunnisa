package main

import "fmt"

func logarithmic(n int) int {
	var result int = 0
	for n > 1 {
		n /= 2
		result += 1
	}
	return result
}

func main() {
	result := logarithmic(100)
	fmt.Println(result)
}

package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n - 1)
	}
}

func main() {
	result := factorial(5) // Change 5 to the desired value for factorial calculation
	fmt.Println(result)
}

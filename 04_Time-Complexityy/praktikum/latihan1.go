package main

import (
	"fmt"
)

func PrimeNumber(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	input1 := 1000000007
	if PrimeNumber(input1) {
		fmt.Println(input1)
		fmt.Println("Bilangan Prima\n")
	} else {
		fmt.Println(input1)
		fmt.Println("Bukan Bilangan Prima\n")
	}

	input2 := 1500450271
	if PrimeNumber(input2) {
		fmt.Println(input2)
		fmt.Println("Bilangan Prima\n")
	} else {
		fmt.Println(input2)
		fmt.Println("Bukan Bilangan Prima\n")
	}
}

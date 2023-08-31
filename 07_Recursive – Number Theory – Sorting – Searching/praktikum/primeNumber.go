package main

import (
	"fmt"
	"math"
)

func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7} // Change this to the numbers you want to check for primality

	for _, number := range numbers {
		prime := checkPrime(number)
		fmt.Println(number, " : ", prime)
	}
}

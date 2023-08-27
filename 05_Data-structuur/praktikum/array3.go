package main

import "fmt"

func main() {
	primes := [5]int{2, 3, 5}

	// cara ke dua

	for i, element := range primes {
		fmt.Println(i, "--->", element)
	}

}

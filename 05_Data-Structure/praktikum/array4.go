package main

import "fmt"

func main() {
	primes := [5]int{2, 3, 5}

	// cara ketiga

	i := 0
	for range primes {
		fmt.Println(primes[i])
		i++
	}

}

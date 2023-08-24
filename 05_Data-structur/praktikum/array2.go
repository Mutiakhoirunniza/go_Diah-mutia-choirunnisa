package main

import "fmt"

func main() {
	primes := [5]int{2, 3, 5}

	// cara pertama//

	for i := 0; i < len(primes); i++ {
		fmt.Println((primes)[i])
	}
}

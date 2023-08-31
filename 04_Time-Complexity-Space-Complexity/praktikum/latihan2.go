package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	halfPow := pow(x, n/2)
	if n%2 == 0 {
		return halfPow * halfPow
	} else {
		return halfPow * halfPow * x
	}
}

func main() {
	x := 2
	n := 3
	fmt.Printf("input : x=%d, n=%d\n", x, n)
	fmt.Println(pow(x, n))

	x = 7
	n = 2
	fmt.Printf("input : x=%d, n=%d\n", x, n)
	fmt.Println(pow(x, n))
}

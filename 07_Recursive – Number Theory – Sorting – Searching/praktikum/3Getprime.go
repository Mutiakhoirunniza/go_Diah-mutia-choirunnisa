package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

func getPrime(position int) int {
	if position <= 0 {
		return -1
	}

	count := 0
	num := 2
	for count < position {
		if isPrime(num) {
			count++
		}
		num++
	}

	return num - 1
}

func main() {
	fmt.Println(getPrime(1))  // 2
	fmt.Println(getPrime(5))  // 11
	fmt.Println(getPrime(8))  // 19
	fmt.Println(getPrime(9))  // 23
	fmt.Println(getPrime(10)) // 29
}

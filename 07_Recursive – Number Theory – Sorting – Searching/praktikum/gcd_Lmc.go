package main

import "fmt"

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a int, b int) int {
	return a * (b / gcd(a, b))
}

func main() {
	number1 := 30 // Replace with the first number
	number2 := 12 // Replace with the second number

	resultGCD := gcd(number1, number2)
	resultLCM := lcm(number1, number2)

	fmt.Printf("%d, %d\n", number1, number2)
	fmt.Printf("GCD: %d\n", resultGCD)
	fmt.Printf("LCM: %d\n", resultLCM)
}

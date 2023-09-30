package main


import "fmt"

func main() {
	a := 10
	b := 5

	// Addition
	addition := a + b
	fmt.Printf("Addition: %d\n", addition)

	// Subtraction
	subtraction := a - b
	fmt.Printf("Subtraction: %d\n", subtraction)

	// Multiplication
	multiplication := a * b
	fmt.Printf("Multiplication: %d\n", multiplication)

	// Division
	if b != 0 {
		division := a / b
		fmt.Printf("Division: %d\n", division)
	} else {
		fmt.Println("Division by zero is not allowed.")
	}
}

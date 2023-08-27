package main

import "fmt"

func getMinMax(numbers ...*int) (min, max int) {
	if len(numbers) == 0 {
		return
	}

	min = *numbers[0]
	max = *numbers[0]

	for _, num := range numbers {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}

	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, a7, a8, min, max int

	fmt.Println("Enter 8 integers:")
	fmt.Scan(&a1, &a2, &a3, &a4, &a5, &a6, &a7, &a8)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6, &a7, &a8)

	

	if max >= 8 {
		fmt.Printf("%d is the maximum number\n", max)
	}
}

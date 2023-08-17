package main

import "fmt"

func main() {
	number := 26

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Println(i)
		}
	}
}

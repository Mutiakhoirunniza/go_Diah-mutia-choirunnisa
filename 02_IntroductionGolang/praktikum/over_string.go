package main

import "fmt"

func main() {
	// LOOPING OVER STRING
	sentence := "Dunia"
	for i := 0; i < len(sentence); i++ {
		fmt.Printf(string(sentence[i]) + "-")
	}
	fmt.Println("   ====== >>>>")
	for pos, char := range sentence {
		fmt.Printf("Character %c start at byte position %d\n", char, pos)
	}
}

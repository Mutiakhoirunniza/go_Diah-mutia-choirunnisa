package main

import (
	"fmt"
)

func shiftLetter(letter rune, offset int) rune {
	ascii := int(letter)
	shifted := (ascii-97+offset)%26 + 97
	return rune(shifted)
}

func shiftString(input string, offset int) string {
	shifted := ""
	for _, char := range input {
		if char != ' ' {
			shifted += string(shiftLetter(char, offset))
		} else {
			shifted += " "
		}
	}
	return shifted
}

func main() {
	var offset int
	var input string

	fmt.Print("offset: ")
	fmt.Scan(&offset)

	fmt.Print("input : ")
	fmt.Scan(&input)

	shiftedString := shiftString(input, offset)
	fmt.Printf("Shifted string: %s\n", shiftedString)
}

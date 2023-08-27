package main

import (
	"fmt"
	"strings"
)

func shiftString(offset int, input string) string {
	var result strings.Builder

	for _, char := range input {
		if char == ' ' {
			result.WriteRune(' ') // Jika karakter adalah spasi, tambahkan spasi ke hasil
		} else {
			shiftedChar := rune(((int(char-'a') + offset) % 26) + int('a'))
			result.WriteRune(shiftedChar)
		}
	}

	return result.String()
}

func main() {
	var offset int
	var input string

	fmt.Print("offset: ")
	fmt.Scan(&offset)

	fmt.Print("input: ")
	fmt.Scan(&input)

	shiftedString := shiftString(offset, input)
	fmt.Printf("Shifted string: %s\n", shiftedString)
}

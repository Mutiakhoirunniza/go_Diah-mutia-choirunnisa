package main

import (
	"fmt"
)

func shiftString(offset int, input string) string {
	shifted := ""

	for _, char := range input {
		if char == ' ' {
			shifted += " "
		} else {
			// Menangani karakter alfabet kecil (a-z)
			shiftedChar := rune(((int(char)-97)+offset)%26 + 97)
			shifted += string(shiftedChar)
		}
	}

	return shifted
}

func main() {
	var offset int
	var input string

	fmt.Print("Masukkan offset: ")
	fmt.Scan(&offset)

	fmt.Print("Masukkan input: ")
	fmt.Scan(&input)

	shifted := shiftString(offset, input)
	fmt.Println("Input:", input)
	fmt.Println("Shifted:", shifted)
}

package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	var result strings.Builder

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			newChar := 'a' + (char-'a'+rune(offset))%26
			result.WriteRune(newChar)
		} else if char >= 'A' && char <= 'Z' {
			newChar := 'A' + (char-'A'+rune(offset))%26
			result.WriteRune(newChar)
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnpj
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}

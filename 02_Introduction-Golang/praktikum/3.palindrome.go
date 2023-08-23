package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Apakah palindrome?")
	fmt.Print("Masukkan kata : ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if isPalindrome(input) {
		fmt.Println("Captured :", input)
		fmt.Println("palindrome.")
	} else {
		fmt.Println("Captured :", input)
		fmt.Println("bukan palindrome.")
	}
}

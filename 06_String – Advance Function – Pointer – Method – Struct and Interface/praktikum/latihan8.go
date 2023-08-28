package main

import (
	"fmt"
	"strings"
)

type Cipher interface {
	Encode(input string) string
	Decode(input string) string
}

type SubstitutionCipher struct {
	substitutions map[rune]rune
}

func NewSubstitutionCipher() *SubstitutionCipher {
	substitutions := map[rune]rune{
		'r': 'i',
		'i': 'r',
		'z': 'a',
		'k': 'p',
		'y': 'b',
	}

	return &SubstitutionCipher{substitutions: substitutions}
}

func (s *SubstitutionCipher) Encode(input string) string {
	var result strings.Builder

	for _, char := range input {
		if newChar, exists := s.substitutions[char]; exists {
			result.WriteRune(newChar)
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func (s *SubstitutionCipher) Decode(input string) string {
	reversedSubstitutions := make(map[rune]rune)
	for key, value := range s.substitutions {
		reversedSubstitutions[value] = key
	}

	var result strings.Builder

	for _, char := range input {
		if originalChar, exists := reversedSubstitutions[char]; exists {
			result.WriteRune(originalChar)
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	cipher := NewSubstitutionCipher()

	var menu int
	fmt.Print("[1] Encrypt\n[2] Decrypt\nChoose your menu? ")
	fmt.Scan(&menu)

	var input string
	fmt.Print("\nInput Student Name: ")
	fmt.Scan(&input)

	if menu == 1 {
		encoded := cipher.Encode(input)
		fmt.Printf("Encode of student’s name %s is %s\n", input, encoded)
	} else if menu == 2 {
		decoded := cipher.Decode(input)
		fmt.Printf("Decode of student’s name %s is %s\n", input, decoded)
	}
}

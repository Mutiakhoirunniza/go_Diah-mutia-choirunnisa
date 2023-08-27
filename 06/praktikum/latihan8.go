package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	for _, char := range s.name {
		if char == ' ' {
			nameEncode += " "
		} else {
			encodedChar := substituteCipher(char, 3)
			nameEncode += string(encodedChar)
		}
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	for _, char := range s.name {
		if char == ' ' {
			nameDecode += " "
		} else {
			decodedChar := substituteCipher(char, -3)
			nameDecode += string(decodedChar)
		}
	}

	return nameDecode
}

func substituteCipher(char rune, shift int) rune {
	shifted := (int(char-'a')+shift+26)%26 + 'a'
	return rune(shifted)
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	fmt.Print("\nInput Student Name: ")
	fmt.Scan(&a.name)

	if menu == 1 {
		encodedName := c.Encode()
		fmt.Printf("\nEncode of student's name %s is : %s", a.name, encodedName)
	} else if menu == 2 {
		decodedName := c.Decode()
		fmt.Printf("\nDecode of student's name %s is : %s", a.name, decodedName)
	}
}

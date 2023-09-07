package main

import "fmt"

func main() {

	func(urutan string) {
		fmt.Println("Hello word" + " bagian " + urutan)
		fmt.Println("alterra academy" + " bagian " + urutan)
	}("3") // tanpa nama

	fmt.Println("Hello word bagian 2")
	fmt.Println("alterra academy bagian 2")
}

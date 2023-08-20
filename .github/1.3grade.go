package main

import "fmt"

func main() {
	var Nilai int = 101

	if Nilai >= 80 && Nilai <= 100 {
		fmt.Println("Grade A")
	} else if Nilai >= 65 && Nilai <= 79 {
		fmt.Println("Grade B")
	} else if Nilai >= 50 && Nilai <= 64 {
		fmt.Println("Grade C")
	} else if Nilai >= 35 && Nilai <= 49 {
		fmt.Println("Grade D")
	} else if Nilai >= 0 && Nilai <= 34 {
		fmt.Println("Grade E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}

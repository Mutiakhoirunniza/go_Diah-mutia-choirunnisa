package main

import "fmt"

func main() {
	// IF, ELSE IF, ELSE

	HOUR := 15
	if HOUR < 12 {
		fmt.Println(" Selamat Pagi ")
	} else if HOUR < 18 {
		fmt.Println(" Selamat Sore ")
	} else {
		fmt.Println(" Selamat Malam ")
	}
}

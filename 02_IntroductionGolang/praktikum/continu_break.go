package main

import "fmt"

func main() {
	// LOOPING WITH CONTUNE & BREAK

	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}
		if i > 3 {
			break
		}
		fmt.Println(i)
	}
}

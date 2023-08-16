package main

import "fmt"

func main() {
	// SWITCH
	var Today int = 3
	switch Today {
	case 1:
		fmt.Println(" Today Is Monday")
	case 2:
		fmt.Println(" Today Is Tuesday ")
	case 3:
		fmt.Println(" Today Is Wednesday ")
	default:
		fmt.Println(" Invalid Date ")
	}

}

package main

import "fmt"

func main() {
	// append & copy

	var colors = []string{"red", "green", "yellow"}
	colors = append(colors, "purple")

	copied_colors := make([]string, 5) //5 adalah cap

	copy(copied_colors, colors)
	fmt.Println(copied_colors)
}

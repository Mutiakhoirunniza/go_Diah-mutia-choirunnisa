package main

import "fmt"

func munculSekali(angka string) []int {
	counts := make(map[int]int)
	var result []int

	for _, digit := range angka {
		num := int(digit - '0')
		counts[num]++
	}

	for _, digit := range angka {
		num := int(digit - '0')
		if counts[num] == 1 {
			result = append(result, num)
			counts[num] = 0
		}
	}

	return result
}

func main() {
	input := "76523752"
	output := munculSekali(input)
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Output: %v\n", output)
}

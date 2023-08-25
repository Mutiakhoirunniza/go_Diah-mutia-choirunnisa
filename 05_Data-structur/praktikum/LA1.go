package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := make(map[string]bool)
	result := []string{}

	for _, name := range arrayA {
		if !merged[name] {
			result = append(result, name)
			merged[name] = true
		}
	}

	for _, name := range arrayB {
		if !merged[name] {
			result = append(result, name)
			merged[name] = true
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}

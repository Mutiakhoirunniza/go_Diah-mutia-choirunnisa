package main

import (
	"fmt"
	"sync"
)

func countFrequency(text string, result chan<- map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()

	frequency := make(map[rune]int)

	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			frequency[char]++
		}
	}

	result <- frequency
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	// Menyiapkan goroutine dan channel
	numWorkers := 4
	resultChannel := make(chan map[rune]int, numWorkers)
	var wg sync.WaitGroup

	// Membagi teks menjadi bagian-bagian
	chunkSize := len(text) / numWorkers
	startIndex := 0
	endIndex := chunkSize

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		if i == numWorkers-1 {
			go countFrequency(text[startIndex:], resultChannel, &wg)
		} else {
			go countFrequency(text[startIndex:endIndex], resultChannel, &wg)
		}
		startIndex = endIndex
		endIndex += chunkSize
	}

	// Mengumpulkan hasil dari goroutine
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Menggabungkan hasil dari goroutine
	totalFrequency := make(map[rune]int)
	for frequency := range resultChannel {
		for char, count := range frequency {
			totalFrequency[char] += count
		}
	}

	// Menampilkan hasil
	for char, count := range totalFrequency {
		fmt.Printf("%c: %d\n", char, count)
	}
}

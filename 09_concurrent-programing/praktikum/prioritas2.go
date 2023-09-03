package main

import (
	"fmt"
	"sync"
)

func MultiplesOfThree(start, end int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if i%3 == 0 {
			ch <- i
		}
	}
}

func main() {
	const start = 1
	const end = 15 // Anda dapat mengubah batas sesuai kebutuhan
	const goroutineCount = 3

	var wg sync.WaitGroup
	ch := make(chan int)

	// Menjalankan goroutine untuk mencari kelipatan 3
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go MultiplesOfThree(start+i, end, ch, &wg)
	}

	// Mengumpulkan hasil dari channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Mencetak hasil dari channel
	for multiple := range ch {
		fmt.Printf("%d adalah kelipatan 3\n", multiple)
	}
}

package main

import (
	"fmt"
	"sync"
)

func printMultiplesOfThree(start, end int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if i%3 == 0 {
			ch <- i
		}
	}
}

func main() {
	const start = 1
	const end = 20
	const goroutineCount = 3
	const bufferSize = 20 // Ubah ukuran buffer menjadi 20

	var wg sync.WaitGroup
	ch := make(chan int, bufferSize)

	// Menjalankan goroutine untuk mencari kelipatan 3
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go printMultiplesOfThree(start+i, end, ch, &wg)
	}

	// Menunggu hingga semua goroutine selesai
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Mencetak hasil dari channel
	for multiple := range ch {
		fmt.Printf("%d adalah kelipatan 3\n", multiple)
	}
}

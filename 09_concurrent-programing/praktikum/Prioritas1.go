package main

import (
	"fmt"
	"sync"
	"time"
)

func printMultiples(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; ; i++ {
		if i%x == 0 {
			fmt.Println(i, "kelipatan dari 10")
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	x := 10
	var wg sync.WaitGroup
	wg.Add(1)

	go printMultiples(x, &wg)

	// Menunggu sampai goroutine selesai
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func printMultiples(x int, max int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= max; i++ {
		if i%x == 0 {
			fmt.Println(i, "kelipatan dari", x)
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	x := 10
	max := 110
	var wg sync.WaitGroup
	wg.Add(1)

	go printMultiples(x, max, &wg)

	// Menunggu sampai goroutine selesai
	wg.Wait()
}

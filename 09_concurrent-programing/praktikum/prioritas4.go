package main

import (
	"fmt"
	"sync"
)

var (
	result int64
	mutex  sync.Mutex
)

func factorial(n int) int64 {
	if n <= 0 {
		return 1
	}
	fact := int64(1)
	for i := 1; i <= n; i++ {
		fact *= int64(i)
	}
	return fact
}

func calculateFactorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Mengunci mutex
	mutex.Lock()
	defer mutex.Unlock()

	result += factorial(n)
}

func main() {
	const numCalculations = 5
	var wg sync.WaitGroup

	// Angka-angka yang akan dihitung faktorialnya
	numbers := []int{5, 6, 7, 8, 9}

	// Mengatur WaitGroup untuk menunggu semua goroutine selesai
	wg.Add(numCalculations)

	// Menjalankan goroutine untuk menghitung faktorial
	for _, n := range numbers {
		go calculateFactorial(n, &wg)
	}

	// Menunggu semua goroutine selesai
	wg.Wait()

	fmt.Printf("Hasil penjumlahan faktorial: %d\n", result)
}

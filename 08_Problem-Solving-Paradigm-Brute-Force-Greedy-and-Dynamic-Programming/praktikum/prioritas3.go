package main

import "fmt"

// Membuat array untuk menyimpan hasil Fibonacci yang sudah dihitung sebelumnya
var memo []int

// Fungsi fibonacci menghitung angka Fibonacci ke-n.
func fibonacci(number int) int {
	// Cek apakah hasil sudah pernah dihitung sebelumnya
	if memo[number] != -1 {
		return memo[number]
	}

	// Jika belum dihitung, hitung dan simpan dalam memo
	if number <= 1 {
		memo[number] = number
	} else {
		memo[number] = fibonacci(number-1) + fibonacci(number-2)
	}

	return memo[number]
}

func main() {
	n := 12 // Ganti dengan nilai n yang diinginkan
	memo = make([]int, n+1)
	for i := range memo {
		memo[i] = -1 // Inisialisasi array memo dengan nilai -1
	}

	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}

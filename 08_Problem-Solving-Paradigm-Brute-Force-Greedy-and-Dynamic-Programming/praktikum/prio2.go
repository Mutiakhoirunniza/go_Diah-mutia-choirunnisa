package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	N := len(jumps)
	if N <= 1 {
		return 0
	}

	// Buat sebuah array untuk menyimpan biaya minimum untuk mencapai setiap batu.
	dp := make([]int, N)

	// Inisialisasi biaya untuk mencapai batu pertama (dp[0] = 0)
	// dan biaya untuk mencapai batu kedua (dp[1]) dengan biaya absolut dari perbedaan tinggi.
	dp[0] = 0
	dp[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	// Iterasi dari batu ketiga hingga batu terakhir untuk menghitung biaya minimum.
	for i := 2; i < N; i++ {
		// Mengambil minimum dari dua opsi:
		// 1. Melompat dari batu sebelumnya (dp[i-1]) dengan biaya absolut perbedaan tinggi.
		// 2. Melompat dari batu dua batu sebelumnya (dp[i-2]) dengan biaya absolut perbedaan tinggi.
		dp[i] = int(math.Min(float64(dp[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1])))), float64(dp[i-2]+int(math.Abs(float64(jumps[i]-jumps[i-2]))))))
	}

	// Hasil biaya minimum yang diperlukan untuk mencapai batu terakhir (N) akan dikembalikan.
	return dp[N-1]
}

func main() {
	// Menggunakan fungsi Frog untuk menguji dengan dua contoh input dan mencetak hasilnya.
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // Output: 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // Output: 40
}

package main

import (
	"fmt"
	"math"
)

func main() {
	// Mendefinisikan matriks 3x3
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	// Inisialisasi variabel untuk menyimpan total diagonal
	diagonalLeftToRight := 0
	diagonalRightToLeft := 0

	// Mendapatkan ukuran matriks
	size := len(matrix)

	// Menghitung total elemen dalam diagonal kiri ke kanan dan diagonal kanan ke kiri
	for i := 0; i < size; i++ {
		diagonalLeftToRight += matrix[i][i]
		diagonalRightToLeft += matrix[i][size-i-1]
	}

	// Menghitung perbedaan mutlak antara total diagonal
	absoluteDifference := int(math.Abs(float64(diagonalLeftToRight - diagonalRightToLeft)))

	// Menampilkan hasil ke konsol dengan keterangan
	fmt.Println("Matriks:")
	for i := 0; i < size; i++ {
		fmt.Println(matrix[i])
	}

	fmt.Printf("Diagonal kiri ke kanan: %d\n", diagonalLeftToRight)
	fmt.Printf("Diagonal kanan ke kiri: %d\n", diagonalRightToLeft)
	fmt.Printf("Perbedaan mutlak antara diagonal: %d\n", absoluteDifference)
}

package main

import "fmt"

func generatePascalPyramid(numRows int) {
	for i := 0; i < numRows; i++ {
		// Menampilkan spasi
		for j := 0; j < numRows-i; j++ {
			fmt.Print(" ")
		}

		// Mengisi piramida dari kiri ke kanan
		for j := 0; j <= i; j++ {
			value := calculatePascalValue(i, j)
			fmt.Printf("%d ", value)
		}

		// Menambahkan baris baru
		fmt.Println()
	}
}

func calculatePascalValue(row, col int) int {
	// Menggunakan rumus kombinasi nCk = n! / (k! * (n-k)!)
	result := 1
	for i := 1; i <= col; i++ {
		result = result * (row - i + 1) / i
	}
	return result
}

func main() {
	numRows := 5
	generatePascalPyramid(numRows)
}

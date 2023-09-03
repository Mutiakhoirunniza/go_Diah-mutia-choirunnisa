package main

import (
	"encoding/json" // Package untuk mengolah JSON
	"fmt"
	"net/http" // Package untuk berinteraksi dengan HTTP
	"sync"     // Package untuk mengelola WaitGroup (sinkronisasi)
)

// Struct Product digunakan untuk merepresentasikan data produk
type Product struct {
	Title    string  `json:"title"`    // Field judul produk
	Price    float64 `json:"price"`    // Field harga produk
	Category string  `json:"category"` // Field kategori produk
}

// Fungsi fetchProducts digunakan untuk mengambil data produk dari URL tertentu dan mengirimkannya ke channel ch
func fetchProducts(url string, wg *sync.WaitGroup, ch chan<- Product) {
	defer wg.Done() // Mengurangkan nilai WaitGroup saat fungsi selesai

	resp, err := http.Get(url) // Mengirim permintaan HTTP GET ke URL
	if err != nil {
		fmt.Println("Error:", err) // Menampilkan pesan kesalahan jika terjadi kesalahan
		return
	}
	defer resp.Body.Close() // Menutup respons HTTP setelah selesai

	var products []Product // Membuat variabel untuk menyimpan produk yang diambil dari respons

	err = json.NewDecoder(resp.Body).Decode(&products) // Mendecode respons JSON ke dalam variabel products
	if err != nil {
		fmt.Println("Error:", err) // Menampilkan pesan kesalahan jika terjadi kesalahan dalam dekoding JSON
		return
	}

	for _, p := range products {
		ch <- p // Mengirim setiap produk ke channel ch
	}
}

func main() {
	url := "https://fakestoreapi.com/products" // URL sumber data produk
	numWorkers := 5                            // Jumlah goroutine yang akan bekerja

	var wg sync.WaitGroup // WaitGroup digunakan untuk nunggu selesai semua goroutine
	wg.Add(numWorkers)    // Menambahkan jumlah goroutine yang harus menyelesaikan pekerjaannya

	ch := make(chan Product) // Membuat channel untuk mengirim data produk

	// Membuat beberapa goroutine untuk mengambil data produk secara paralel
	for i := 0; i < numWorkers; i++ {
		go fetchProducts(url, &wg, ch)
	}

	go func() {
		wg.Wait() // Menunggu semua goroutine selesai
		close(ch) // Menutup channel setelah semua pekerjaan selesai
	}()

	fmt.Println("Product data")

	// Membaca data produk dari channel ch dan mencetaknya
	for p := range ch {
		fmt.Printf("title: %s\nprice: %.2f\ncategory: %s\n===\n", p.Title, p.Price, p.Category)
	}
}

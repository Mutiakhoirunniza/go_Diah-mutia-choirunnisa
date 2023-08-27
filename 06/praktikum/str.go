package main

import (
	"fmt"
	"strings"
)

func main() {
	value := "cat;dog"
	substring := value[5:len(value)] // Mengambil substring dari indeks 5 hingga akhir
	fmt.Println(substring)

	// Mengganti karakter "a" ke "i" hanya 4 kali
	s := "bakatakatakataka"
	t := strings.Replace(s, "a", "i", 4)
	fmt.Println(t)

	// insert

	p := "purpel"
	index := 4
	q := p[:index] + "UNGU" + p[index:]
	fmt.Println(p, q)

}

package main

import (
	"fmt"
	"strings"
)

const (
	Str    = "bodyright"
	Substr = "right"
)

func main() {
	// 1. panjang string
	kalimat := "areudoing"
	panjangKalimat := len(kalimat)
	fmt.Println(panjangKalimat) // menghitung jumlah karakter dalam string

	// 2. membandingkan string
	Str1 := "something"
	Str2 := "someone"

	fmt.Println(Str1 == Str2) // membandingkan Str1 dan Str2

	// 3. mengandung
	res := strings.Contains(Str, Substr)
	fmt.Println(res) // mengecek apakah Substr ada di dalam Str
}

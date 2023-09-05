package main

import "fmt"

type Kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (mobil *Mobil) Berjalan() {
	mobil.TambahKecepatan(10)
}

func (mobil *Mobil) TambahKecepatan(kecepatanBaru int) {
	mobil.kecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLamban := Mobil{}
	mobilLamban.Berjalan()

	fmt.Printf("Kecepatan Mobil Cepat: %d km/jam\n", mobilCepat.kecepatanPerJam)
	fmt.Printf("Kecepatan Mobil Lamban: %d km/jam\n", mobilLamban.kecepatanPerJam)
}

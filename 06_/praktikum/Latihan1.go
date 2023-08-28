package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Car struct {
	tipeMobil  string
	bahanBakar float64
	factor     float64
}

func (c Car) estimasiJarak() float64 {
	distance := c.bahanBakar / c.factor
	return distance
}

func main() {
	var carType string
	fmt.Print("Masukkan jenis mobil : ")
	fmt.Scanln(&carType)
	carType = strings.ToLower(carType)

	var fuelInput string
	fmt.Print("Masukkan jumlah bahan bakar (dalam liter): ")
	fmt.Scanln(&fuelInput)

	bahanBakar, err := strconv.ParseFloat(fuelInput, 64)
	if err != nil {
		fmt.Println("Input bahan bakar tidak valid.")
		os.Exit(1)
	}

	var factor float64
	if carType == "suv" {
		factor = 1.5
	} else {
		fmt.Println("Jenis mobil tidak valid.")
		os.Exit(1)
	}

	myCar := Car{
		tipeMobil:  carType,s
		bahanBakar: bahanBakar,
		factor:     factor,
	}

	estimasiJarak := myCar.estimasiJarak()

	fmt.Printf("Estimasi jarak yang bisa ditempuh: %.0f\n", estimasiJarak)
}

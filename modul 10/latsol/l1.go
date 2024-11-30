package main

import "fmt"

func main() {
	var berat, beratSisa, beratKg, biayaPerKg, totalBiaya int

	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&berat)

	beratKg = berat / 1000
	beratSisa = berat % 1000
	biayaPerKg = 10000
	totalBiaya = biayaPerKg * beratKg

	if beratKg > 10 {
		beratSisa = 0
	} else if beratSisa >= 500 {
		totalBiaya += beratSisa * 5
	} else {
		totalBiaya += beratSisa * 15
	}

	fmt.Printf("Berat parsel: %d gram\n", berat)
	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, beratSisa)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKg*beratKg, beratSisa*5)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}

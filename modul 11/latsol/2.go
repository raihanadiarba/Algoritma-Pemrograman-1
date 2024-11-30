package main

import "fmt"

func main() {
	var durasi, tarif int
	var tipe string

	fmt.Print("Jenis kendaraan (Motor/Mobil/Truk) : ")
	fmt.Scan(&tipe)
	fmt.Print("Durasi parkir (dalam jam) : ")
	fmt.Scan(&durasi)

	switch {
	case tipe == "mobil":
		tarif = 5000 * durasi
	case tipe == "motor":
		tarif = 2000 * durasi
	case tipe == "truk":
		tarif = 8000 * durasi
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Println("Rp", tarif)
}

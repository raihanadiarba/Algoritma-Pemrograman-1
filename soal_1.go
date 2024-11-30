package main

import "fmt"

func main() {
	var panjang, lebar_alas, tinggi float64
	var volume, luas_permukaan float64

	fmt.Scan(&panjang, &lebar_alas, &tinggi)

	volume = panjang * lebar_alas * tinggi

	luas_permukaan = 2*(panjang*lebar_alas) + 2*(panjang*tinggi) + 2*(lebar_alas*tinggi)

	fmt.Print("\n")
	fmt.Printf("%.1f\n", volume)
	fmt.Printf("%.1f\n", luas_permukaan)
}

package main

import "fmt"

type Hitung struct {
	Panjang  float64
	Luas     float64
	Lebar    float64
	Keliling float64
}

func main() {
	var t Hitung

	// Input dari user
	fmt.Print("Panjang: ")
	fmt.Scanln(&t.Panjang)
	fmt.Print("Lebar: ")
	fmt.Scanln(&t.Lebar)

	// Menghitung
	t.Keliling = 2 * (float64(t.Panjang) + float64(t.Lebar))
	t.Luas = float64(t.Panjang) * float64(t.Lebar)

	// Output
	fmt.Println("\nHasil Hitung:")
	fmt.Printf("Luas : %.1f\n", t.Luas)
	fmt.Printf("Keliling : %.1f\n", t.Keliling)
}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var daerahA, daerahB, daerahC, daerahD int
	var totalTitik int

	fmt.Print("Masukkan jumlah titik hujan: ")
	fmt.Scan(&totalTitik)

	for i := 0; i < totalTitik; i++ {
		x := rand.Float64()
		y := rand.Float64()

		// Cek posisi titik
		if x < 0.5 && y < 0.5 {
			daerahA++
		} else if x >= 0.5 && y < 0.5 {
			daerahB++
		} else if x >= 0.5 && y >= 0.5 {
			daerahC++
		} else {
			daerahD++
		}
	}

	curahA := float64(daerahA) * 0.0001
	curahB := float64(daerahB) * 0.0001
	curahC := float64(daerahC) * 0.0001
	curahD := float64(daerahD) * 0.0001

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", curahA)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", curahB)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", curahC)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", curahD)
}

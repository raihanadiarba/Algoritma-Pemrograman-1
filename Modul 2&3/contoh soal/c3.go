package main

import "fmt"

func main() {
	var bb, tb, bmi float64
	fmt.Print("masukkan berat badan (kg): ")
	fmt.Scan(&bb)
	fmt.Print("masukkan tinggi badan (m): ")
	fmt.Scan(&tb)
	bmi = bb / (tb * tb)
	fmt.Printf("bmi anda  : %.2f\n", bmi)
}

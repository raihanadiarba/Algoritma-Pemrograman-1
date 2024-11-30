package main

import (
	"fmt"
)

func main() {
	var fx float64
	fmt.Print("Enter f(x): ")
	fmt.Scan(&fx)

	if fx <= 5 {
		fmt.Println("Nilai f(x) harus lebih besar dari 5 untuk menghindari kesalahan matematika.")
		return
	}

	x := 2/(fx-5) - 5
	fmt.Printf("Nilai x adalah: %.0f\n", x) // Output x sebagai bilangan bulat

}

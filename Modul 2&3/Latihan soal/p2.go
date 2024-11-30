package main

import "fmt"

func main() {
	var volume, luas, r float64
	fmt.Print("Jari Jari = ")
	fmt.Scan(&r)
	luas = 4.0 * 3.1415926535 * r * r
	volume = (4.0 / 3.0) * 3.1415926535 * r * r * r
	fmt.Printf("Bola dengan jejari %0.f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}

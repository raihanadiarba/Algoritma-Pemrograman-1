package main

import "fmt"

func main() {

	var angka, xy float64

	fmt.Scan(&angka)
	xy = float64(int(angka)) + 1

	for selesai := false; !selesai; {
		angka += 0.1
		angka = float64(int(angka*10+0.5)) / 10
		if angka >= xy {
			fmt.Printf("%.0f\n", angka)
			selesai = true
		} else {
			fmt.Printf("%.1f\n", angka)
		}

	}

}

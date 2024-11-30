package main

import "fmt"

func main() {
	var bunga, pita string
	var jumlahBunga int

	for {
		fmt.Printf("Bunga %d: ", jumlahBunga+1)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		if jumlahBunga == 0 {
			pita = bunga
		} else {
			pita = pita + " - " + bunga
		}
		jumlahBunga++
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", jumlahBunga)
}

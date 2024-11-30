package main

import "fmt"

func main() {
	fmt.Println("=== Program Validasi Tahun (if-else) ===")

	var tahun int
	fmt.Print("Masukkan tahun yang ingin dicek: ")
	fmt.Scan(&tahun)

	if tahun <= 0 {
		fmt.Println("Error: Tahun tidak valid!")
	} else {
		if tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0) {
			fmt.Printf("Tahun %d adalah tahun kabisat\n", tahun)
			fmt.Println("Jumlah hari: 366 hari")
		} else {
			fmt.Printf("Tahun %d adalah tahun biasa\n", tahun)
			fmt.Println("Jumlah hari: 365 hari")
		}
	}
}

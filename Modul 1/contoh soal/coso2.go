package main

import "fmt"

func main() {
	uang := 98700
	hargaPerBarang := 14500

	jumlahBarang := uang / hargaPerBarang
	sisaUang := uang % hargaPerBarang

	fmt.Printf("Pelanggan bisa membeli %d barang.\n", jumlahBarang)
	fmt.Printf("Sisa uang setelah membeli barang: %d rupiah.\n", sisaUang)
}

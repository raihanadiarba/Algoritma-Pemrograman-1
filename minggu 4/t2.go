package main

import "fmt"

// Struct untuk menyimpan data transaksi
type Transaksi struct {
	NamaBarang  string
	Jumlah      int
	HargaSatuan float64
	TotalHarga  float64
}

func main() {
	var t Transaksi

	// Input dari user
	fmt.Print("Masukkan Nama Barang: ")
	fmt.Scanln(&t.NamaBarang)
	fmt.Print("Masukkan Jumlah: ")
	fmt.Scanln(&t.Jumlah)
	fmt.Print("Masukkan Harga Satuan: Rp ")
	fmt.Scanln(&t.HargaSatuan)

	// Menghitung total harga
	t.TotalHarga = float64(t.Jumlah) * t.HargaSatuan

	// Output
	fmt.Println("\nInformasi Transaksi:")
	fmt.Println("Nama Barang:", t.NamaBarang)
	fmt.Println("Jumlah:", t.Jumlah)
	fmt.Printf("Harga Satuan: Rp %.2f\n", t.HargaSatuan)
	fmt.Printf("Total Harga: Rp %.2f\n", t.TotalHarga)
}

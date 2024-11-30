package main

import "fmt"

type Karyawan struct {
	Nama      string
	GajiPokok float64
	Tunjangan float64
	Potongan  float64
	TotalGaji float64
}

func main() {
	var k Karyawan

	// Input data karyawan
	fmt.Print("Masukkan Nama Karyawan: ")
	fmt.Scanln(&k.Nama)
	fmt.Print("Masukkan Gaji Pokok (Rp): ")
	fmt.Scanln(&k.GajiPokok)
	fmt.Print("Masukkan Tunjangan (Rp): ")
	fmt.Scanln(&k.Tunjangan)
	fmt.Print("Masukkan Potongan (Rp): ")
	fmt.Scanln(&k.Potongan)

	k.TotalGaji = (k.GajiPokok + k.Tunjangan) - k.Potongan

	fmt.Printf("\nInformasi Karyawan:\n")
	fmt.Printf("Nama: %s\n", k.Nama)
	fmt.Printf("Gaji Pokok: Rp %.2f\n", k.GajiPokok)
	fmt.Printf("Tunjangan: Rp %.2f\n", k.Tunjangan)
	fmt.Printf("Potongan: Rp %.2f\n", k.Potongan)
	fmt.Printf("Total Gaji: Rp %.2f\n", k.TotalGaji)
}

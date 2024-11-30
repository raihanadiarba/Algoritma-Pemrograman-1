package main

import "fmt"

func main() {
	var nama, nim, kelas, prodi string
	fmt.Print("masukkan nama: ")
	fmt.Scan(&nama)
	fmt.Print("masukkan nim: ")
	fmt.Scan(&nim)
	fmt.Print("masukkan kelas: ")
	fmt.Scan(&kelas)
	fmt.Print("prodi: ")
	fmt.Scan(&prodi)

	fmt.Println("Perkenalkan saya adalah ", nama, "salah satu mahasiswa Prodi ", prodi, "dari kelas", kelas, "dengan NIM", nim)

}

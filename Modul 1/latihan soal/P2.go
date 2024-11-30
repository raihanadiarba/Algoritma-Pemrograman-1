package main

import "fmt"

func main() {
	var satu, dua, tiga string
	fmt.Print("Nama : ")
	fmt.Scan(&satu)
	fmt.Print("NIM : ")
	fmt.Scan(&dua)
	fmt.Print("Kelas : ")
	fmt.Scan(&tiga)
	fmt.Println("Perkenalkan saya adalah " + satu + " " + "salah satu mahasiswa Prodi S1-IF dari kelas " + tiga + " " + "dengan NIM " + dua)

}

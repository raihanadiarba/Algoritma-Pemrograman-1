package main

import "fmt"

func main() {
	var angka int
	fmt.Print("masukkan angka : ")
	fmt.Scan(&angka)
	fmt.Println(angka%2 == 0)
}

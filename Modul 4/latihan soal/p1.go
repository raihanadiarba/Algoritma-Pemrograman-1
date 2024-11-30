package main

import "fmt"

func main() {
	var hargawal, diskon, total int
	fmt.Scan(&hargawal)
	fmt.Scan(&diskon)

	total = hargawal - (hargawal*diskon)/100
	fmt.Println(total)
}

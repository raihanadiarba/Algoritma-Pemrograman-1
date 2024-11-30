package main

import "fmt"

func main() {
	var luas, keliling, p, l int
	fmt.Println("=> Mencari keliling dan luas dari persegi Panjang")
	fmt.Print("Panjang = ")
	fmt.Scan(&p)
	fmt.Print("Lebar = ")
	fmt.Scan(&l)
	//rumus
	luas = p * l
	keliling = 2*p + 2*l
	fmt.Println("Keliling = ", keliling, "\nLuas =", luas)
}

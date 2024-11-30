package main

import "fmt"

func main() {
	var a, b, c, d, e, hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)
	hasil = a + b + c + d + e
	fmt.Println("Hail Penjumlahan ", hasil)
}

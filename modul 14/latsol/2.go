package main

import "fmt"

func main() {
	var hasil string
	var n int

	fmt.Scan(&n)
	hasil = "prima"

	if n%n == 0 && n%2 != 0 {
		fmt.Println(hasil)
	} else {
		fmt.Println("Bukan Prima")
	}
}

package main

import "fmt"

func main() {
	var v1, v2, nn, hasil int
	fmt.Scan(&v1, &v2)
	for nn = 1; nn <= v2; nn++ {
		hasil += hasil + v1

	}
	fmt.Println(hasil)
}

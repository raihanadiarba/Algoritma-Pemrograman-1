package main

import "fmt"

func main() {
	var x, y int
	var hasil = 1
	fmt.Scan(&x, &y)
	for i := 0; i < y; i++ {
		hasil *= x
	}
	fmt.Println(hasil)
}

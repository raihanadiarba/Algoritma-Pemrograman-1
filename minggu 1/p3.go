package main

import "fmt"

func main() {
	var k, l, r, pi float64
	pi = 3.14
	fmt.Println("=> Menghitung luas dan keliling lingkaran")
	fmt.Print("r = ")
	fmt.Scan(&r)
	//rumus
	k = pi * r * r
	l = 2 * pi * r
	fmt.Println("Keliling = ", k)
	fmt.Println("Luas = ", l)

}

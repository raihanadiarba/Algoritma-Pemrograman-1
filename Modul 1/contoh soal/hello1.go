package main


import "fmt"

func main() {
	//variable
	var k, l, r, pi float64
	pi = 3.14
	//input
	fmt.Scan(&r)
	//rumus
	k = pi * r * r
	l = 2 * pi * r
	//outputhasil
	fmt.Println(k)
	fmt.Println(l)

}

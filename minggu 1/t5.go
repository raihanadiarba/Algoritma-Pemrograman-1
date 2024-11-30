package main

import "fmt"

func main() {
	var hb1, hb2, hb3 float64
	fmt.Scan(&hb1, &hb2, &hb3)
	hj1 := hb1 * 1.05
	hj2 := hb2 * 1.05
	hj3 := hb3 * 1.05

	println(hj1, hj2, hj3)

}

package main

import "fmt"

func main() {
	var X int
	fmt.Print("x = ")
	fmt.Scan(&X)
	d1 := X / 100
	d2 := (X % 100) / 10
	d3 := X % 10
	fmt.Println(d1, d2, d3)
}

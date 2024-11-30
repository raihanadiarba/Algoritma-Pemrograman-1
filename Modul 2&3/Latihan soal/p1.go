package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Print("Enter x: ")
	fmt.Scan(&x)

	fx := 2/(x+5) + 5
	fmt.Println(fx)
}

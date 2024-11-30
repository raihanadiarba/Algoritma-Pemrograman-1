package main

import "fmt"

func main() {
	var bmi, tb, bb float64
	fmt.Scan(&bmi, &tb)

	bb = bmi * tb * tb

	fmt.Printf("%.f\n", bb)
}

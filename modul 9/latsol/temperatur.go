package main

import "fmt"

func main() {
	var a, b, c, d, e float64
	fmt.Scan(&a, &b, &c, &d, &e)
	if a >= b && b >= c && c >= d && d >= e {
		fmt.Println("temperatur stabil turun")
	} else if a <= b && b <= c && c <= d && d <= e {
		fmt.Println("temperatur stabil naik")
	} else {
		fmt.Println("temperatur tidak stabil")
	}
}

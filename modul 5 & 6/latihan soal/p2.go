package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var vol, j, t float64
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&j, &t)
		vol = (1.0 / 3.0) * (math.Pi * j * j * t)
		fmt.Println(vol)

	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3, y1, y2, y3 float64
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)

	pjgAB := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	pjgBC := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	pjgCA := math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2))

	plgpjg := math.Max(pjgAB, math.Max(pjgBC, pjgCA))

	fmt.Println(plgpjg)
}

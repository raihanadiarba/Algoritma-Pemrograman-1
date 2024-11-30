package main

import "fmt"

func main() {
	var b, x int
	var y bool
	fmt.Scan(&b, &x)

	for b := x % 10; y && x > 0; b = x % 10 {
		x %= 10
		if x/10 <= b {
			y = true
		}
	}

	fmt.Println(y)
}

package main

import "fmt"

func main() {
	var x, y int
	var end bool
	fmt.Scan(&x, &y)
	for end = false; !end; {
		x = x - y
		fmt.Println(x)
		end = x <= 0
	}
	fmt.Println(x == 0)
}

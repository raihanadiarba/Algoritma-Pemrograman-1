package main

import "fmt"

func main() {
	var bilangan, digit int

	digit = 0
	fmt.Scan(&bilangan)

	for bilangan > 0 {
		bilangan = bilangan / 10
		digit++
	}
	fmt.Println(digit)
}

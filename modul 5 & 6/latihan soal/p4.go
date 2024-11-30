package main

import "fmt"

func main() {
	var n int
	var faktorial = 1
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		faktorial *= i
	}
	fmt.Println(faktorial)
}

package main

import "fmt"

func main() {
	var a, b int
	var j int
	fmt.Scan(&a, &b)
	for j = a; j <= b; j += 1 {
		fmt.Println(j, " ")
	}
}

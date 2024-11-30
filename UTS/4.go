package main

import "fmt"

func main() {
	var input, out int

	for fmt.Scan(&input); input != 0; fmt.Scan(&input) {
		out += input * input
	}

	fmt.Println(out)
}

package main

import "fmt"

func main() {
	for i := 1; i <= 3; i += 1 {
		for j := 0; j < 3; j += 1 {
			fmt.Printf("(%d%d)", i, j)

		}
	}
	fmt.Println()
}

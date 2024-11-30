package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	for j := 1; j < bilangan; j += 1 {
		if j%2 != 0 {
			fmt.Print(j, " ")
		}

	}
}

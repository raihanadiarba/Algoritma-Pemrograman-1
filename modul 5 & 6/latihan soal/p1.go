package main

import "fmt"

func main() {
	var X, hasil int
	fmt.Scan(&X)

	for i := 1; i <= X; i++ {
		hasil += i
	}

	fmt.Printf("%d\n", hasil)

}

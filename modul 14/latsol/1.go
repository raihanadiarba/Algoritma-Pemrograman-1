package main

import "fmt"

func main() {
	var n, i, jumlahGanjil int

	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		if i%2 != 0 {
			jumlahGanjil++
		}
	}

	fmt.Printf("Terdapat %d bilangan ganjil\n", jumlahGanjil)
}

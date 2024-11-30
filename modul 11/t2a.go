package main

import "fmt"

func main() {
	var N, literasi int
	literasi = 0
	fmt.Scan(&N)
	for literasi != N {
		literasi = literasi + 1
		fmt.Println(literasi)
	}
}

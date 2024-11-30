package main

import "fmt"

func main() {
	var N, literasi int
	literasi = 0
	fmt.Scan(&N)
	for {
		literasi++
		fmt.Println(literasi)
		if literasi == N {
			break
		}
	}
}

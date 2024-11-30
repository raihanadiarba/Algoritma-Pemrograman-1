package main

import "fmt"

func main() {
	var tinggi int
	fmt.Scan(&tinggi)
	i := 1
	for i := 1; i <= tinggi; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

package main

import "fmt"

func main() {
	var tinggi int
	fmt.Scan(&tinggi)
	
	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= tinggi-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

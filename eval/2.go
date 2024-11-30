package main

import "fmt"

func main() {
	var tinggi int
	fmt.Scan(&tinggi)

	for i := tinggi; i >= 1; i-- {
		for j := 1; j <= tinggi-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("    ")

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

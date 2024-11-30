package main

import "fmt"

func main() {
	var tinggi, i, j int

	fmt.Scan(&tinggi)

	for i = 1; i <= tinggi; i++ {
		for j = 1; j <= tinggi-i; j++ {
			fmt.Print("  ")
		}
		for j = 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("  ")
		for j = 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	for i = tinggi - 1; i >= 1; i-- {
		for j = 1; j <= tinggi-i; j++ {
			fmt.Print("  ")
		}
		for j = 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("  ")
		for j = 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

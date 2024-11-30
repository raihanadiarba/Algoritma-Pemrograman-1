package main

import "fmt"

func main() {
	var num int
	var isPrime bool = true

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num)

	if num <= 1 {
		isPrime = false
	} else {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Printf("%d adalah bilangan prima.\n", num)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", num)
	}
}

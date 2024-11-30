package main

import "fmt"

func main() {
	var number int
	var cloop bool
	for cloop = true; cloop; {
		fmt.Scan(&number)
		cloop = number <= 0
	}
	fmt.Printf("%d adalah bilangan positif\n", number)
}

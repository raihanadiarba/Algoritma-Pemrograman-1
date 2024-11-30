package main

import "fmt"

func main() {
	var word string
	var ulang int
	fmt.Scan(&word, &ulang)
	counter := 0
	for done := false; !done; {
		fmt.Println(word)
		counter++
		done = (counter >= ulang)
	}
}

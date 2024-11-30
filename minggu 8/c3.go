package main

import "fmt"

func main() {
	var kar, x int
	fmt.Scan(&kar)

	if kar%7 == 0 {
		x = kar / 1

	} else {
		x = (kar / 7) + 1
	}
	fmt.Println(x)

}

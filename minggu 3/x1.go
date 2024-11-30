package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	ganjil := bilangan%2 != 0
	fmt.Print(ganjil)
}

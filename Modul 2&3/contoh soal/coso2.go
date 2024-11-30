package main

import "fmt"

func main() {
	var alas, luas, tinggi float64
	fmt.Print("Alas : ")
	fmt.Scan(&alas)
	fmt.Print("tinggi : ")
	fmt.Scan(&tinggi)
	luas = (alas * tinggi / 2)
	fmt.Println(luas)
}

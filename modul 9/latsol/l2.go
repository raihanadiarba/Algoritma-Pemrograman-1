package main

import "fmt"

func main() {
	var n int
	var txt string
	fmt.Scan(&n)
	txt = "bukan"
	if n%2 == 0 && n < 0 {
		txt = "genap negatif"
	}
	fmt.Println(txt)
}

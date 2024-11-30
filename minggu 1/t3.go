package main

import "fmt"

func main() {
	println("=> Fungsi F(x,y)")
	var hitung, x, y float64
	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Print("y = ")
	fmt.Scan(&y)
	hitung = 1/(3*x*x+10) + 10*y + 7
	fmt.Println(hitung)

}

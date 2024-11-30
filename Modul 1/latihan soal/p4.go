package main

import "fmt"

func main() {
	var Fahrenheit, C float64
	fmt.Scan(&Fahrenheit)
	// Rumus Celcius ke Fahrenheit: °C = 5/9 x (°F - 32)
	C = (Fahrenheit - 32) * 5 / 9
	fmt.Println(C)
}

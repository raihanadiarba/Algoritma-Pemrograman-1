package main

import "fmt"

func main() {
	var x, y, i, temp int

	fmt.Scan(&x, &y)

	i = 0
	temp = x
	for temp >= y {
		temp -= y
		i++
	}
	fmt.Println(i)
}

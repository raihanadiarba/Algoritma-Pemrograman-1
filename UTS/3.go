package main

import "fmt"

func main() {
	var input, berat int
	fmt.Scan(&input)

	blm := 0
	sdh := 0

	for i := 0; i < input; i++ {
		fmt.Scan(&berat)

		blm += berat
		sdh += berat * 90 / 100

	}
	fmt.Println(blm, sdh)
}

package main

import "fmt"

func main() {
	var pH float64
	fmt.Scan(&pH)

	switch {
	case pH < 0 || pH > 14:
		fmt.Println("Input tidak valid, rentang pH 0-14")
	case pH >= 6.5 && pH <= 8.6:
		fmt.Println("Air Layak Minum")
	case pH < 6.5 || pH > 8.6:
		fmt.Println("Air Tidak Layak Minum")
	}
}
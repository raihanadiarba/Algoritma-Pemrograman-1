package main

import "fmt"

func main() {
	var tak, ikk float64
	fmt.Scan(&tak)
	if tak >= 3.00 {
		fmt.Scan(&ikk)
		var kategori string
		if ikk < 2.00 {
			kategori = "Poor"
		} else if 2.00 <= ikk && ikk <= 2.75 {
			kategori = "Fair"
		} else if 2.76 <= ikk && ikk <= 3.00 {
			kategori = "Sastifactory"
		} else if 3.01 <= ikk && ikk <= 3.50 {
			kategori = "Very Good"
		} else {
			kategori = "Excellent"
		}
		fmt.Printf("cumlaude dengan IPK %.2f dan status predikat TAK adalah %s\n", tak, kategori)
	} else {
		fmt.Println("Tidak Cumlaude")
	}
}

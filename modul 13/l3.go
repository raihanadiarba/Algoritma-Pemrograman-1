package main

import "fmt"

func main() {
	var target, donasi, total, donatur int
	fmt.Scan(&target)
	donatur = 0
	total = 0

	for {
		fmt.Scan(&donasi)
		total += donasi
		donatur++
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donatur, donasi, total)

		if total >= target {
			break
		}
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, donatur)
}

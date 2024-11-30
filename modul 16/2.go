package main

import "fmt"

func main() {
	var x, input string
	var n, count, firstPos int
	var found bool

	fmt.Print("Masukkan string yang dicari: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	count = 0
	firstPos = -1
	found = false

	for i := 1; i <= n; i++ {
		fmt.Printf("Masukkan string ke-%d: ", i)
		fmt.Scan(&input)

		if input == x {
			count++
			if !found {
				firstPos = i
				found = true
			}
		}
	}
	fmt.Println("\nHasil:")

	if found {
		fmt.Println("a. Ya, string ditemukan")
	} else {
		fmt.Println("a. Tidak, string tidak ditemukan")
	}

	if firstPos != -1 {
		fmt.Printf("b. Ditemukan pada posisi ke-%d\n", firstPos)
	} else {
		fmt.Println("b. String tidak ditemukan")
	}

	fmt.Printf("c. String muncul sebanyak %d kali\n", count)

	if count >= 2 {
		fmt.Println("d. Ya, terdapat minimal dua string yang sama")
	} else {
		fmt.Println("d. Tidak, tidak terdapat minimal dua string yang sama")
	}
}

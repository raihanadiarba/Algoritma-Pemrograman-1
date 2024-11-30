// Raihan Adi Arba_103112400071
// 103112400071
package main

import "fmt"

func main() {
	var n, A, B, C, AB, BA_103112400071 int
	var bol bool = true
	fmt.Scan(&n, &A, &B)
	AB = A + B
	BA_103112400071 = B - A

	if n < 3 {
		fmt.Println("Deret tidak valid")
		return
	}

	for i := 3; i <= n; i++ {
		fmt.Scan(&C)
		AB += C

		if C-B != BA_103112400071 {
			bol = true
		}
		B = C
	}

	if bol {
		fmt.Printf("Jumlah total: %d\n", AB)
	} else {
		fmt.Println("Bukan deret aritmatika")
	}
}

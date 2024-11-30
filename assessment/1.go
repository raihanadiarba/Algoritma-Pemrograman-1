// Raihan Adi Arba
// 103112400071
package main

import "fmt"

func main() {
	var sisiA_103112400071, sisiB, sisiC int
	fmt.Scan(&sisiA_103112400071, &sisiB, &sisiC)

	if sisiA_103112400071 >= sisiB+sisiC || sisiB >= sisiA_103112400071+sisiC || sisiC >= sisiA_103112400071+sisiB {
		fmt.Println("bukan segitiga")
		return
	}

	if sisiA_103112400071 == sisiB && sisiB == sisiC {
		fmt.Println("Segitiga sama sisi")
	} else if sisiA_103112400071 == sisiB || sisiB == sisiC || sisiA_103112400071 == sisiC {
		fmt.Println("Segitiga sama kaki")
	} else if sisiA_103112400071*sisiA_103112400071 == sisiB*sisiB+sisiC*sisiC || sisiB*sisiB == sisiA_103112400071*sisiA_103112400071+sisiC*sisiC || sisiC*sisiC == sisiA_103112400071*sisiA_103112400071+sisiB*sisiB {
		fmt.Println("Segitiga siku-siku")
	} else {
		fmt.Println("Segitiga sembarang")
	}

}

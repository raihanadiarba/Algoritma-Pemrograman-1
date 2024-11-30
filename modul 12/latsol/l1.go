package main

import "fmt"

func main() {
	var user, pass string
	var failed int

	failed = 0
	fmt.Scan(&user, &pass)
	for user != "Admin" || pass != "Admin" {
		fmt.Scan(&user, &pass)
		failed++
	}
	fmt.Printf("%d Percobaan gagal login\n", failed)
}

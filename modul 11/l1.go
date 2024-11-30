package main

import "fmt"

func main() {
	var user, pass string
	failedAttempts := 0
	isLoggedIn := false

	for !isLoggedIn {
		fmt.Scanln(&user, &pass)

		if user == "admin" && pass == "admin" {
			isLoggedIn = true
			fmt.Printf("%d\nLogin berhasil\n", failedAttempts)
		} else {
			failedAttempts++
		}
	}
}

package main

import "fmt"

func main() {
	var huruf rune
	var konsonan bool
	fmt.Scanf("%c", &huruf)
	konsonan = (huruf >= 'a' && huruf <= 'z' || huruf >= 'A' && huruf <= 'Z') &&
		!(huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' ||
			huruf == 'A' || huruf == 'I' || huruf == 'U' || huruf == 'E' || huruf == 'O')
	fmt.Print(konsonan)
}

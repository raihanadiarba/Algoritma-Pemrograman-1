package main

import "fmt"

func main() {
	var digit int

	fmt.Scan(&digit)

	satu := digit % 10
	dua := (digit / 10) % 10
	tiga := (digit / 100) % 10
	empat := (digit / 1000) % 10
	lima := digit / 10000

	balik := satu*10000 + dua*1000 + tiga*100 + empat*10 + lima

	fmt.Println(balik)
}

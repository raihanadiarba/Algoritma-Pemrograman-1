package main

import "fmt"

const PI float64 = 3.14

type tabung struct {
	tinggi, jari          int
	luas, volume          float64
	luasAlas, luasDinding float64
}

func main() {
	var t tabung
	fmt.Scan(&t.tinggi, &t.jari)

	t.luasAlas = PI * float64(t.jari) * float64(t.jari)
	t.luasDinding = float64(t.tinggi) * (2 * PI * float64(t.jari))
	t.luas = (2 * t.luasAlas) + t.luasDinding
	t.volume = t.luasAlas * float64(t.tinggi)

	fmt.Printf("Luas: %.2f\n", t.luas)
	fmt.Printf("Volume: %.2f\n", t.volume)
}

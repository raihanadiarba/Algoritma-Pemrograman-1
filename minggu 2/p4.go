package main

import "fmt"

type Hitung struct {
	tcelcius, dfahrenheit, dreamur, dkelvin, hkelvin, hreamur, hcelsius, hfahrenheit float64
}

func main() {
	var t Hitung
	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&t.tcelcius)

	// Celsius -> Fahrenheit
	t.hfahrenheit = (t.tcelcius * 9 / 5) + 32

	// Celsius -> Reamur
	t.hreamur = t.tcelcius * 4 / 5

	// Celsius -> Kelvin
	t.hkelvin = t.tcelcius + 273.15

	fmt.Printf("Temperatur Celsius : %.0f\n", t.tcelcius)
	fmt.Printf("Derajat Reamur : %.0f\n", t.hreamur)
	fmt.Printf("Derajat Fahrenheit : %.0f\n", t.hfahrenheit)
	fmt.Printf("Derajat Kelvin : %.0f\n", t.hkelvin)

}

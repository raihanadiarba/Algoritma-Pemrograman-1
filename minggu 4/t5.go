package main

import "fmt"

type Hitung struct {
	HargaBeli, HargaJual, JmlSaham, Keliling, totalInvest, Penjualan, KeuntunganBersih, KeuntunganKotor, BiayaTransaksi, PajakKeuntungan float64
}

func main() {
	var t Hitung

	// Input dari user
	fmt.Print("Harga Beli: Rp ")
	fmt.Scanln(&t.HargaBeli)
	fmt.Print("Harga Jual: Rp ")
	fmt.Scanln(&t.HargaJual)
	fmt.Print("Jumlah Saham : Rp ")
	fmt.Scanln(&t.JmlSaham)

	// Menghitung
	t.totalInvest = float64(t.HargaBeli) * float64(t.JmlSaham)
	t.Penjualan = float64(t.HargaJual) * float64(t.JmlSaham)
	t.KeuntunganKotor = float64(t.Penjualan) - float64(t.totalInvest)
	t.BiayaTransaksi = float64(t.Penjualan) * 0.002
	t.PajakKeuntungan = float64(t.KeuntunganKotor) * 0.10

	if t.KeuntunganKotor > 0 {
		t.PajakKeuntungan = t.KeuntunganKotor * 0.10
	} else {
		t.PajakKeuntungan = 0
	}
	t.KeuntunganBersih = float64(t.KeuntunganKotor) - float64(t.BiayaTransaksi) - float64(t.PajakKeuntungan)

	// Output
	fmt.Println("\nHasil Hitung: ")
	fmt.Printf("Total Investasi Awal : Rp %.2f\n", t.totalInvest)
	fmt.Printf("Total Penjualan : Rp %.2f\n", t.Penjualan)
	fmt.Printf("Keuntungan Kotor : Rp %.2f\n", t.KeuntunganKotor)
	fmt.Printf("Biaya Transaksi : Rp %.2f\n", t.BiayaTransaksi)
	fmt.Printf("Pajak Keuntungan : Rp %.2f\n", t.PajakKeuntungan)
	fmt.Printf("Keuntungan Bersih : Rp %.2f\n", t.KeuntunganBersih)

}

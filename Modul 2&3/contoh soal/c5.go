package main

import "fmt"

//Harga beli: hb
//Harga jual: hj
//Jumlah saham: js
//Total investasi awal: tia
//Total penjualan: tp
//Keuntungan Kotor: kk
//Biaya transaksi: bt
//Pajak keuntungan: pk
//Keuntungan bersih: kb

type Informasi_Investasi_Saham struct {
	hb, hj, js, tia, tp, kk, bt, pk, kb float64
}

func main() {
	var info Informasi_Investasi_Saham
	fmt.Println("Informasi Investasi Saham: ")
	fmt.Print("Harga beli: Rp ")
	fmt.Scan(&info.hb)
	fmt.Print("Harga jual: Rp ")
	fmt.Scan(&info.hj)
	fmt.Print("Jumlah Saham: ")
	fmt.Scan(&info.js)

	info.tia = info.hb * info.js
	info.tp = info.hj * info.js
	info.kk = info.tp - info.tia
	info.bt = 0.2 * info.tp / 100
	info.pk = 10 * info.kk / 100
	info.kb = info.kk - info.bt - info.pk

	fmt.Printf("Total Investasi: Rp %.2f\n", info.tia)
	fmt.Printf("Total Penjualan: Rp %.2f\n", info.tp)
	fmt.Printf("Keuntungan Kotor: Rp %.2f\n", info.kk)
	fmt.Printf("Biaya Transaksi: Rp %.2f\n", info.bt)
	fmt.Printf("Pajak Keuntungan: Rp %.2f\n", info.pk)
	fmt.Printf("Keuntungan Bersih: Rp %.2f\n", info.kb)
}
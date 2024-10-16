package main

import "fmt"

func main() {
	var belnja_awal, besar_diskon int
	fmt.Scan(&belnja_awal, &besar_diskon)
	fmt.Scanln()
	belanja_akhir := belnja_awal - (belnja_awal * besar_diskon/100)
	fmt.Print(belanja_akhir)
}

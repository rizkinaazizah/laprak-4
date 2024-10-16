package main

import "fmt"

func main() {
	var bmi, tb float64

	fmt.Print("Masukkan nilai BMI:")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan (meter):")
	fmt.Scan(&tb)
	bb := (tb*tb) * bmi
	fmt.Printf("Berat badan:%.f",bb)
}

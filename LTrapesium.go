package main

import "fmt"

func main() {
	var alasAtas, alasBawah, tinggi float64

	fmt.Print("Masukkan panjang alas atas trapesium: ")
	fmt.Scanln(&alasAtas)

	fmt.Print("Masukkan panjang alas bawah trapesium: ")
	fmt.Scanln(&alasBawah)

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&tinggi)

	luas := 0.5 * (alasAtas + alasBawah) * tinggi
	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}

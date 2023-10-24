package main

import "fmt"

// Struct Car yang memiliki properti type dan fuelIn
type Car struct {
	Type   string
	FuelIn float64 // Bahan bakar dalam liter
}

// Method untuk menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar
func (c Car) EstimateDistance() float64 {
	// Konversi bahan bakar dari liter ke mil (1.5 L / mil)
	// Perkiraan jarak = bahan bakar (dalam liter) / konversi (1.5 L / mil)
	distance := c.FuelIn / 1.5
	return distance
}

func main() {
	// Membuat instance dari struct Car
	myCar := Car{
		Type:   "Sedan",
		FuelIn: 30.0, // 30 liter bahan bakar
	}

	// Memanggil method EstimateDistance
	estimatedDistance := myCar.EstimateDistance()

	fmt.Printf("Mobil tipe %s dengan %0.2f liter bahan bakar dapat ditempuh sejauh %0.2f mil.\n", myCar.Type, myCar.FuelIn, estimatedDistance)
}

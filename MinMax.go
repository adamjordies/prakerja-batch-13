package main

import "fmt"

// Function untuk mencari nilai maksimum dan minimum dari array
func findMaxMin(numbers []int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}

	max, min := numbers[0], numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

func main() {
	var numbers [6]int

	fmt.Println("Masukkan 6 angka:")
	for i := 0; i < 6; i++ {
		fmt.Printf("Angka ke-%d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	max, min := findMaxMin(numbers[:])

	fmt.Printf("Nilai maksimum: %d\n", max)
	fmt.Printf("Nilai minimum: %d\n", min)
}

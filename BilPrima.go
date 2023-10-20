package main

import (
	"fmt"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	} else if number <= 3 {
		return true
	} else if number%2 == 0 || number%3 == 0 {
		return false
	}
	i := 5
	for i*i <= number {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	var num int
	fmt.Print("Masukkan bilangan: ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("Input bukan bilangan bulat.")
		return
	}

	if isPrime(num) {
		fmt.Printf("%d adalah bilangan prima.\n", num)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", num)
	}
}

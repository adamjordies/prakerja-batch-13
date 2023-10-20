package main

import (
	"fmt"
)

func isMultipleOf7(number int) bool {
	return number%7 == 0
}

func main() {
	var num int
	fmt.Print("Masukkan bilangan: ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("Input bukan bilangan bulat.")
		return
	}

	if isMultipleOf7(num) {
		fmt.Printf("%d adalah kelipatan 7.\n", num)
	} else {
		fmt.Printf("%d bukan kelipatan 7.\n", num)
	}
}

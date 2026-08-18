package main

import "fmt"

func main() {
	const usdToEur = 0.92
	usdAmount := 0.0

	fmt.Println("--- USD to EUR Conversion Chart ---")
	for i := 10.0; i <= 100; i += 10.0 {
		usdAmount = i * usdToEur

		fmt.Printf("$%.2f USD = €%.2f EUR\n", i, usdAmount)
	}
}

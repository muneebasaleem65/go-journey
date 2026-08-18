package main

import "fmt"

func calculateTotal(numbers []float64) float64 {
	total := 0.0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	prices := []float64{19.99, 5.50, 42.00}

	result := calculateTotal(prices)

	fmt.Println("Total: ", result)
}

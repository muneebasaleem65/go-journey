package main

import "fmt"

func applyDiscounts(startingTotal float64, listOfDiscounts ...float64) float64 {

	for _, value := range listOfDiscounts {
		startingTotal -= value
	}

	return startingTotal
}

func main() {
	starting := 150.00

	listOfNums := []float64{10.50, 5.00, 25.00}

	result := applyDiscounts(starting, listOfNums...)

	fmt.Println("Total", result)
}

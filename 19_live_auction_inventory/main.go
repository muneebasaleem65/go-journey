package main

import "fmt"

func main() {
	auctionMap := map[int]float64{101: 150.50, 102: 45.00}

	var maxVal float64
	var highestItem int

	for k, v := range auctionMap {
		fmt.Printf("Item #%d currently has a high bid of $%.2f\n", k, v)
		if v > maxVal {
			maxVal = v
			highestItem = k
		}
	}

	fmt.Printf("The winning item so far is: Item #%d\n", highestItem)
}

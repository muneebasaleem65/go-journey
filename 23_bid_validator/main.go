package main

import "fmt"

type Calculator func(float64) bool

func processBid(itemID int, bidAmount float64, validator Calculator) {
	if validator(bidAmount) {
		fmt.Println("Bid Accepted for item #", itemID)
	} else {
		fmt.Println("Bid Rejected for item # ", itemID)
	}
}

func main() {
	minimumBidRule := func(a float64) bool {
		return a > 50.0
	}

	processBid(101, 10.0, minimumBidRule)
}

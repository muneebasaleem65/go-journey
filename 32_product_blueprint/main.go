package main

import "fmt"

type auctionItem struct {
	id            string
	name          string
	currentPrice  float64
	highestBidder string
}

func newAuctionItem(id string, name string, currentPrice float64) *auctionItem {
	return &auctionItem{
		id:            id,
		name:          name,
		currentPrice:  currentPrice,
		highestBidder: "No Bids yet",
	}
}

func (a *auctionItem) placeBid(bidderName string, bidAmount float64) {
	if bidAmount > a.currentPrice {
		a.highestBidder = bidderName
		a.currentPrice = bidAmount

		fmt.Printf("Success! New high bid of $%.2f by %s on %s\n", bidAmount, bidderName, a.name)
	} else {
		fmt.Printf("Bid rejected! $%.2f is too low for %s\n", bidAmount, a.name)
	}
}

func main() {
	item := newAuctionItem("101", "Vintage Prada", 500.00)
	fmt.Println("Initial Item State", item)

	fmt.Println("\n--- Simulation starting ---")

	item.placeBid("Muneeba", 450.00)

	item.placeBid("Wikki", 650.00)

	fmt.Println("\nFinal Item State", item)
}

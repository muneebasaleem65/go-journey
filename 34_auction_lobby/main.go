package main

import "fmt"

type auctionItem struct {
	id, name, highestBidder string
	currentPrice            float64
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
	inventory := make(map[string]*auctionItem)

	inventory["101"] = &auctionItem{id: "101", name: "vintage watch", currentPrice: 150.0}
	inventory["102"] = &auctionItem{id: "102", name: "retro console", currentPrice: 150.0}

	inventory["101"].placeBid("Muneeba", 250.00)
	inventory["102"].placeBid("Muneeba", 270.00)

	for _, val := range inventory {
		fmt.Println(val.name, val.highestBidder, val.currentPrice)
	}
}

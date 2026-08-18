package main

import "fmt"

type logger interface {
	logAction(action string)
}

// creating an empty struct
type consoleLogger struct{}

func (c consoleLogger) logAction(action string) {
	fmt.Println("[Console log]" + action)
}

type dbLogger struct {
	databaseName string
}

func (d dbLogger) logAction(action string) {
	fmt.Println("Saved to database " + d.databaseName)
}

type auctionItem struct {
	name         string
	currentPrice float64
}

func (a *auctionItem) updatePrice(newPrice float64, logEngine logger) {
	a.currentPrice = newPrice

	logEngine.logAction("\nPrice updated for " + a.name)
}

func main() {
	consoleLog := consoleLogger{}

	dbLog := dbLogger{databaseName: "Auction_Production_DB"}

	inventory := &auctionItem{name: "Vintage Rolex", currentPrice: 500.0}

	inventory.updatePrice(700.0, dbLog)

	inventory.updatePrice(650.0, consoleLog)
}

package main

import "fmt"

type product struct {
	name  string
	stock int
}

func checkLowStock(products []product) {
	for _, item := range products {
		if item.stock < 6 {
			fmt.Println(item.name, "Stock low")
		} else {
			fmt.Println(item.name, "Stock ok")
		}
	}
}

func main() {
	inventory := []product{
		{name: "Laptop", stock: 5},
		{name: "Phone", stock: 12},
		{name: "Monitor", stock: 3},
	}

	checkLowStock(inventory)
}

package main

import "fmt"

type shippable interface {
	getShippingCost() float64
}

type book struct {
	title  string
	weight float64
}

func (b book) getShippingCost() float64 {
	return b.weight * 1.5
}

type electronics struct {
	name  string
	price float64
}

func (e electronics) getShippingCost() float64 {
	if e.price < 100 {
		return 15.0
	}
	return 0.0
}

//with a standalone function

// func calculateTotalShipping(cart []shippable) float64 {
// 	total := 0.0

// 	for _, item := range cart {
// 		total += item.getShippingCost()
// 	}

// 	return total
// }

func main() {
	//declaring a slice using interface name

	cart := []shippable{
		book{title: "Go blueprints", weight: 2.5},
		electronics{name: "Gaming mouse", price: 45.00},
	}

	//call the standalone function and pass the slice

	// grandTotal := calculateTotalShipping(cart)
	// fmt.Printf("Grand total shipping: $%.2f\n", grandTotal)

	totalShipping := 0.0

	for _, item := range cart {
		totalShipping += item.getShippingCost()
	}

	fmt.Printf("Total Shipping Cost: $%.2f\n", totalShipping)

}

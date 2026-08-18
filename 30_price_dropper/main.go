package main

import "fmt"

func applyDiscount(itemPrice *float64) {
	*itemPrice -= 10.0
}

func main() {
	itemPrice := 45.50

	applyDiscount(&itemPrice)

	fmt.Println(itemPrice)
}

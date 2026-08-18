package main

import "fmt"

func main() {
	const criticallevel = 10

	for stock := 24; stock >= 0; stock -= 3 {
		if stock <= criticallevel {
			fmt.Println("Warning! Stock low:", stock)
		} else {
			fmt.Println("Stock OK:", stock)
		}
	}
}
package main

import "fmt"

func main() {
	items := []float64{10.50, 4.99, 25.00, 7.25}

	total := 0.0

	for i := 0; i < len(items); i++ {
		total = total + items[i]
	}
	fmt.Println(total)
}

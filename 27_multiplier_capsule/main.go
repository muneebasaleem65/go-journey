package main

import "fmt"

func multiplier() func() int {
	count := 1

	return func() int {
		count = count * 2

		return count
	}
}

func main() {
	multiplying := multiplier()

	fmt.Println(multiplying())
	fmt.Println(multiplying())
	fmt.Println(multiplying())
	fmt.Println(multiplying())
}

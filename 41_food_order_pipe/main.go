package main

import "fmt"

func prepChef(pipe chan string) {
	pipe <- "Cheeseburger"
	pipe <- "Chicken Burger"
	pipe <- "Veggie Burger"
	close(pipe)
}

func main() {
	burgerSlide := make(chan string, 3)

	go prepChef(burgerSlide)

	for burger := range burgerSlide {
		fmt.Println("[Cashier] Handed [" , burger ,"] to customer!")
	}
}

package main

import "fmt"

func main() {
	for num := 1; num <= 5; num++ {
		switch num {
		case 1, 2:
			var name string
			if num == 1 {
				name = "Soda & Popcorn"
			}
			if num == 2 {
				name = "Juice & Nachos"
			}
			fmt.Println("Combo #", num, name, "is a Drink Special ($8.50)")
		case 3:
			var name string
			name = "Water & candy"
			fmt.Println("Combo #", num, name, "is a Solo Deal ($6.00)")
		case 4, 5:
			var name string
			if num == 4 {
				name = "Large Bucket & 2 Drinks"
			}
			if num == 5 {
				name = "Mega Box & 4 Drinks"
			}
			fmt.Println("Combo #", num, name, "is a Family Pack ($18.00)")
		}
	}
}

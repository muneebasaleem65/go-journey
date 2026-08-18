package main

import (
	"fmt"
)

func main() {
	for dayNum := 1; dayNum <= 7; dayNum++ {
		switch dayNum {
		case 1, 2, 3, 4:
			var name string
			if dayNum == 1 {
				name = "Monday"
			}
			if dayNum == 2 {
				name = "Tuesday"
			}
			if dayNum == 3 {
				name = "Wednesday"
			}
			if dayNum == 4 {
				name = "Thursday"
			}

			fmt.Printf("%s is a regular workday.\n", name)
		case 5:
			fmt.Println("Friday! The weekend is almost here")
		case 6, 7:
			var name string
			if dayNum == 6 {
				name = "Saturday"
			}
			if dayNum == 7 {
				name = "Sunday"
			}
			fmt.Printf("%s is the weekend! Time to relax.\n", name)
		}
	}
}

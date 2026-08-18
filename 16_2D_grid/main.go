package main

import "fmt"

func main() {
	vals := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//change the very center which is number 5 vals[rows][columns]
	vals[1][1] = 11

	fmt.Println("Center value is now:", vals[1][1])
	fmt.Println("Full Grid", vals)
}

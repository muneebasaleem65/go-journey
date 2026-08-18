package main

import "fmt"

func main() {
	lightColor := "green"

	switch lightColor {
	case "red":
		fmt.Println("Stop completely!")
	case "yellow":
		fmt.Println("Slow down and prepare to stop.")
	case "green":
		fmt.Println("Go safely.")
	default:
		fmt.Println("Traffic light is broken! Proceed with caution.")
	}
}

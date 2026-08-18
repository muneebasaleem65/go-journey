package main

import "fmt"

func main() {
	const pi = 3.14159

	fmt.Println("Circle Area Calculator:")
	for radius := 5.0; radius <= 10.0; radius++ {
		area := pi * radius * radius
		fmt.Printf("Radius: %v | Area: %v\n", radius, area)
	}
}

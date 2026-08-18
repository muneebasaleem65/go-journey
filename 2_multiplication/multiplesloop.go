package main

import "fmt"

func main() {
	fmt.Println("Table of 2")
	for i := 2; i <= 20; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("Table of 3")
	for i := 1; i <= 10; i++ {
		fmt.Println(i * 3)
	}
}

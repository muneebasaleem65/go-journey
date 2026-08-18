package main

import "fmt"

func main() {
	for i := 10; i >= 1; i-- {
		if i==5{
			fmt.Println(i," - Halfway there!!")
		} else{
			fmt.Println(i)
		}
	}
}
package main

import "fmt"

func displayAddressAndValue(ptr *int) {
	//print the pointer variable itself to see the raw memor address

	fmt.Println("Raw memory address location", ptr)

	fmt.Println("The value hiding inside that address", *ptr)

}

func main() {
	num := 42

	fmt.Println("pointer Inspection")
	//passing address of 'num' into the function
	displayAddressAndValue(&num)
}

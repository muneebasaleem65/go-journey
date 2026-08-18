package main

import "fmt"

//this function expects an int slice so if we want to print any type of slice or provided types of slice from this function
// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//we will use 'any' type or provide all expected types or write 'comparable' and create a generic fuction

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T int | string | bool] struct {
	elements []T
}

func main() {
	nums := []int{1,2,3,4,5}
	printSlice(nums)
	names := []string{"John doe", "Alex"}
	printSlice(names)

	myStack := stack[int]{
		elements: []int{1,2,3},
	}

	fmt.Println(myStack)
}

package main

import "fmt"

//type Calculator func(int) int

//func runMathRule(number int,mathRule Calculator){}
//is same as below

func runMathRule(number int, mathRule func(int) int) {
	result := mathRule(number)

	fmt.Println("This machine processed the number. Result:", result)
}

func main() {
	//creating anonymouse function inside main function as go does not allow named functions inside the main function
	 
	doubleRule := func(a int) int {
		return a * 2
	}

	addTenRule := func(a int) int {
		return a + 10
	}

	runMathRule(5, doubleRule)
	runMathRule(5, addTenRule)
}

package main

import "fmt"

func main(){
	for age :=5; age<=70; age++{
		if age < 6 {
			fmt.Printf("Age %d: Free ($0.00)\n", age)

		} else if age >= 6 && age <= 18 {
			fmt.Printf("Age %d: Child Ticket ($12.50)\n", age)

		} else if age >= 65 {
			fmt.Printf("Age %d: Senior Ticket ($15.00)\n", age)

		} else {
			fmt.Printf("Age %d: Adult Ticket ($25.00)\n", age)
		}
	}
}

package main

import "fmt"

func main() {
	employees := map[string]string{"Alice": "Engineering", "Bob": "Design", "Charlie": "Sales"}

	//add new employee
	employees["Diana"] = "Marketing"
	//delete bob
	delete(employees, "Bob")
	//Print alice's department
	fmt.Println(employees["Alice"])
	//Print the entire map
	fmt.Println(employees)
}

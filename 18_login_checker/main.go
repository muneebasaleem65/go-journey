package main

import "fmt"

func main() {
	activeUsers := map[string]bool{"alex": true, "sam": false}

	fmt.Println(activeUsers)

	_, ok := activeUsers["john"]

	if ok {
		fmt.Println("User exists")
	} else {
		fmt.Println("User doesn't exist")
	}
}

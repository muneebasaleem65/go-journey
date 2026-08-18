package main

import "fmt"

func changeName(name *string) {
	*name = "Pro_Go_Developer"
	fmt.Println("In changename", *name)
}

func main() {
	name := "Beginner"

	changeName(&name)

	fmt.Println("After changename", name)
}

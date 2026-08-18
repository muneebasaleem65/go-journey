package main

import "fmt"

func main() {
	var scores = []int{45, 82, 94, 61, 75}

	scores = append(scores, 88, 99)

	fmt.Println(len(scores))
	for i := 0; i < len(scores); i++ {
		if scores[i] >=80{
			fmt.Println(scores[i])
		}
	}

}

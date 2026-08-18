package main

import "fmt"

func createLog(listItems ...string) string {

	fullSentence := ""

	for _, item := range listItems {
		fullSentence = fullSentence + item + " "
	}

	return fullSentence
}

func main() {
	listOfWords := []string{"User", "Alex", "placed", "a", "bid", "of", "$150"}

	log := createLog(listOfWords...)

	fmt.Print(log)
}

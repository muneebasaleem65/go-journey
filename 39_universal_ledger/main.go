package main

import "fmt"

func summarizeLedger[T int | string](entries []T) int {
	for _, entry := range entries {
		fmt.Println("[Log Entry]: ", entry)
	}
	return len(entries)
}

func main() {
	summarizeLedger([]int{101,102,103,104})
	summarizeLedger([]string{"IOS101","MS104"})


}

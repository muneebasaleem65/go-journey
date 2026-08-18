package main

import "fmt"

func trackUserAction(userName string) func() {
	trackingCount := 0

	return func() {
		trackingCount++
		fmt.Println(userName, "performed action #", trackingCount)
	}
}

func main() {
	alexTracker := trackUserAction("Alex")

	alexTracker()
	alexTracker()
}

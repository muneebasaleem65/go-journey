package main

import "fmt"

// type orderStatus int

// const (
// 	Received orderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type orderStatus string

const (
	Received  orderStatus = "received"
	Confirmed orderStatus = "confirmed"
	Prepared  orderStatus = "prepared"
	Delivered orderStatus = "delivered"
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeOrderStatus(Confirmed)
}

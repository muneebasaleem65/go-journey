package main

import "fmt"

type notifier interface {
	send(message string)
}

type email struct {
	address string
}

func (e email) send(message string) {
	fmt.Println("Email sent to", e.address, ":", message)
}

type sms struct {
	phoneNumber string
}

func (s sms) send(message string) {
	fmt.Println("SMS sent to", s.phoneNumber, ":", message)
}

func notification(msg string, n notifier) {
	n.send(msg)
}

func main() {
	viaEmail := email{address: "muneeba.centosquare@gmail.com"}
	viaSms := sms{phoneNumber: "03927367464"}

	notification("Sending via email", viaEmail)
	notification("Sending via sms", viaSms)

}

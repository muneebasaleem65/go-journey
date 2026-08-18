package main

import "fmt"

//designing the interface(contract)
type paymentGateway interface {
	//an inerface doesn't store data fields, 
	//it stores list of function promises
	pay(amount float64) bool
}

//we don't need to use the 'implements' keyword
//for a class to use the interface like in php

//just build the struct and attach amethod that
//macthes the contract 

//building the struct
type stripe struct{
	apiKey string
}

//creating the method that matches the contract
func (s stripe) pay(amount float64) bool{
	fmt.Printf("Processing $%.2f safely through stripe API..\n", amount)
	return true
}

//building another struct
type payPal struct{
	userEmail string
}

func (p payPal) pay(amount float64) bool{
	fmt.Printf("Processing $%.2f instantly via paypal Gateway..\n", amount)
	return true
}

//interface is useful as now we only have to write
//a single checkout func that accepts contract as an
//input param.It doesn't care if the user picked stripe or paypal

func checkoutProcessor(amount float64, gateway paymentGateway){
	fmt.Println("Checkout Initialized")
	gateway.pay(amount)
}

func main(){
	stripeProcessor := stripe{apiKey: "sk_live_123"}
	payPalProcessor := payPal{userEmail: "alex@example.com"}

	//pass stripe into the processor
	checkoutProcessor(99.50, stripeProcessor)
	//pass paypal in the same processor
	checkoutProcessor(45.50,payPalProcessor)
}
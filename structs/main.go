package main

import (
	"fmt"
	"time"
)

type customer struct{
	name string
}

//struct syntax reads as
//type structName structkeyword{}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

//create a method to change status in struct
//func (params holds o as in convention structname) functionName(value that needs to be changed type of the value that s being changed){}

// receiver type
func (o *order) changeStatus(status string) {
	//no need to dereference *o.status as structs does it automatically
	o.status = status
}

func main() {
	//creating struct instance
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}

	myOrder.changeStatus("confirmed")

	//add a value in the instance
	myOrder.createdAt = time.Now()

	//get a specific data
	fmt.Println(myOrder.status)

	//create muliple instances
	myOrder2 := order{
		id:        "2",
		amount:    100,
		status:    "delivered",
		createdAt: time.Now(),
		customer: customer{
			name: "John",
		},
	}

	//get order using the funcion

	myOrder3 := newOrder("3", 500, "confirmed")

	fmt.Println("Order struct", myOrder, myOrder2, myOrder3)

	//to write the structs that are used only one time
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}

package main

import "fmt"

type serverState int

const (
	Offline serverState = iota
	Provisioning
	Active
	Terminated
)

func bootServer(currentState serverState) {
	switch currentState {
	case Offline:
		fmt.Println("Starting server virtualization...")
	case Active:
		fmt.Println("Server is healthy and accepting traffic.")
	}
}

func main() {
	bootServer(Active)
}

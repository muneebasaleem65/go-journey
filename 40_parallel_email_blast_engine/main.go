package main

import (
	"fmt"
	"sync"
)

func sendEmail(recipient string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Worker] email sent to ", recipient)
}

func main() {
	var wg sync.WaitGroup

	emailAddresses := []string{"alex@test.com", "sam@test.com", "john@test.com"}

	for i := 0; i < len(emailAddresses); i++ {
		wg.Add(1)
		go sendEmail(emailAddresses[i], &wg)
	}

	wg.Wait()
	fmt.Println("All workers are done sending concurrent email")
}

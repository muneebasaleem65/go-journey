package main

import (
	"fmt"
	"sync"
)

type trafficTracker struct {
	mu    sync.Mutex
	views map[string]int
}

func newTrafficTracker() *trafficTracker {
	return &trafficTracker{
		views: make(map[string]int),
	}
}

func (t *trafficTracker) incrementView(urlPath string){
	t.mu.Lock()
	defer t.mu.Unlock()
	t.views[urlPath]++
}

func main() {
	tracker := newTrafficTracker()

	var wg sync.WaitGroup

	fmt.Println("--- Launching 1,000 Concurrent Traffic Hit Workers ---")

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(){
			defer wg.Done()

			tracker.incrementView("/home")
		}()
	}

	wg.Wait()

	fmt.Println("--- All Background Operations Complete ---")
	fmt.Printf("Total views for '/home': %d\n", tracker.views["/home"])
}

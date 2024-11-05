package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	count int
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func incrementCounter(numIncrements int) {
	defer wg.Done()
	for i := 0; i < numIncrements; i++ {
		mutex.Lock()
		count++
		mutex.Unlock()
	}
}

func decrementCounter(numDecrements int) {
	defer wg.Done()
	for i := 0; i < numDecrements; i++ {
		mutex.Lock()
		count--
		mutex.Unlock()
	}
}

func logStatus(ctx context.Context, interval time.Duration) {
	defer wg.Done()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Status logging stopped.")
			return
		case <-ticker.C:
			mutex.Lock()
			log.Printf("Current count: %d\n", count)
			mutex.Unlock()
		}
	}
}

func main() {
	count = 0
	numIncrements := 500000
	numDecrements := 300000
	statusInterval := 100 * time.Millisecond
	timeout := 5 * time.Second

	// Create a context with a timeout to gracefully handle shutdown
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Start incrementing and decrementing in goroutines
	wg.Add(3)
	go incrementCounter(numIncrements)
	go decrementCounter(numDecrements)
	go logStatus(ctx, statusInterval)

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Printf("Final count: %d\n", count)
}

// import (
// 	"fmt"
// 	"sync"
// )

// var count = 0
// var wg sync.WaitGroup
// var mu sync.Mutex

// func inc() {
// 	mu.Lock()
// 	count++
// 	mu.Unlock()
// 	wg.Done()
// }

// func start() {
// 	for i := 0; i < 1000000; i++ {
// 		wg.Add(1)
// 		go inc()
// 	}
// }

// func main() {
// 	count = 0
// 	start()
// 	wg.Wait()
// 	fmt.Println(count)
// }

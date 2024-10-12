package main

import (
	"fmt"
	"sync"
)

var count = 0
var wg sync.WaitGroup
var mu sync.Mutex

func inc() {
	mu.Lock()
	count++
	mu.Unlock()
	wg.Done()
}

func start() {
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go inc()
	}
}

func main() {
	count = 0
	start()
	wg.Wait()
	fmt.Println(count)
}

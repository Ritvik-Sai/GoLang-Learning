package main

import (
	"fmt"
	"sync"
	"time"
)

func ShowMutex() {
	var mu sync.Mutex
	counter := 0

	for i := 1; i <= 5; i++ {
		go func(id int) {
		mu.Lock()
			defer mu.Unlock()
		  counter++
			fmt.Println("goroutine", id, "increase counter to", counter)
		}(i) }

	time.Sleep(time.Second) // wait for goroutines
	fmt.Println("final counter:", counter)
}

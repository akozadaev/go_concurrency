package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	// Горутина-потребитель
	go func() {
		mu.Lock()
		for !ready {
			cond.Wait() // разблокирует mu и ждёт Signal/Broadcast
		}
		fmt.Println("Work is ready!")
		mu.Unlock()
	}()

	// Горутина-производитель
	time.Sleep(1 * time.Second)
	mu.Lock()
	ready = true
	cond.Broadcast() // или Signal()
	mu.Unlock()

	time.Sleep(100 * time.Millisecond)
}

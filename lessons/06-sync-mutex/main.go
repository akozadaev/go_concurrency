package main

import (
	"fmt"
	"sync"
)

func main() {
	//  go run -race main.go
	var mu sync.Mutex
	var wg sync.WaitGroup

	mapA := map[string]int{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			mapA["one"] += 1
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(mapA)
}

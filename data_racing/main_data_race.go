package main

import (
	"fmt"
	"sync"
)

func main() {
	// go run -race main.go
	var wg sync.WaitGroup

	mapA := map[string]int{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mapA["one"] += 1
		}()
	}

	wg.Wait()

	fmt.Println(mapA)
}

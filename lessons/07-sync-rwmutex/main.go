package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var config = map[string]string{"key": "value"}

func updateConfig(key, value string) {
	mu.Lock()
	defer mu.Unlock()
	config[key] = value
}

func readConfig(key string) string {
	mu.RLock()
	defer mu.RUnlock()
	return config[key]
}
func main() {
	var wg = sync.WaitGroup{}

	// читатели
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Millisecond)
			fmt.Printf("Прочитали %s \n", readConfig("key"))
		}()
	}

	// писатель
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond)
		updateConfig("key", "value2")
	}()

	wg.Wait()
}

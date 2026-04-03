package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initDB() {
	fmt.Println("Connecting to DB...")
	// имитация подключения
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initDB)
			fmt.Println("i:", i)
		}()
	}
	wg.Wait()
	fmt.Println("All done")
}

package main

import (
	"bytes"
	"fmt"
	"sync"
)

// bufPool переиспользует буферы между горутинами.
// Не кладите в Pool объекты с секретами без полного обнуления содержимого.
var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			b := bufPool.Get().(*bytes.Buffer)
			b.Reset()
			defer bufPool.Put(b)
			_, _ = fmt.Fprintf(b, "worker-%d", id)
			fmt.Println(b.String())
		}(i)
	}
	wg.Wait()
}

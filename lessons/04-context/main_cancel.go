package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Получен сигнал отмены: ", ctx.Err())
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Работаю")
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

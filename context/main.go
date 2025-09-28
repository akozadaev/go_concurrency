package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err() // ошибка таймаута
	case <-time.After(2 * time.Second):
		return "успешное завершение", nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second) // таймаут
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // Успех
	time.Sleep(500 * time.Millisecond)
	cancel()
	defer cancel()
	result, err := task(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("Превышен лимит ожидания")
		} else {
			fmt.Printf("Ошибка: %v\n", err)
		}
		//return
	}
	fmt.Println(result)

}

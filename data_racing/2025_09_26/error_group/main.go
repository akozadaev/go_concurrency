package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(100 * time.Millisecond)
		return errors.New("Задача 1 не выполнена")
	})

	g.Go(func() error {
		select {
		case <-time.After(200 * time.Millisecond):
			fmt.Println("Задача 2 выполнена успешно")
			return nil
		case <-ctx.Done():
			fmt.Println("Задача 2 отменена")
			return errors.New("Задача 11 не выполнена")
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	}
}

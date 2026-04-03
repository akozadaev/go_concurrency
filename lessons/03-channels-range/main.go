package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i * 10
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for val := range ch {
		fmt.Printf(" %d", val)
	}

}

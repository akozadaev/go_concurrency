package main

import (
	"fmt"
	"time"
)

func main() {
	go numbers()
	go alpabet()

	//time.Sleep(2 * time.Second)
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alpabet() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

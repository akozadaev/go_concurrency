package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 2)
	go numbers(done)
	go alpabet(done)

	<-done
	<-done
}

func numbers(done chan bool) {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
	done <- true
}

func alpabet(done chan bool) {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
	done <- true
}

//1a23b4c5de
//1a23b4c5de

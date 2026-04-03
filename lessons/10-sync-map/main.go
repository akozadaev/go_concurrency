package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("name", "Alex")
	m.Store("age", 45)

	if v, ok := m.Load("name"); ok {
		fmt.Println("Name:", v)
	}

	m.Range(func(key, value interface{}) bool {
		fmt.Println("Key:", key, "Value:", value)
		return true
	})
	
	m.Delete("age")

	if v, ok := m.Load("name"); ok {
		fmt.Println("Name:", v)
	}

}

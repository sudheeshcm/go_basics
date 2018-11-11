package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup:", r)
	}
	wg.Done()
}

func say(s string) {
	defer cleanup() // usually put at top of the surrounding function
	// method that make sure the final stuffs are done even in panic
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear... a 2") // Raising a panic error
		}
	}
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("there")
	wg.Wait()
}

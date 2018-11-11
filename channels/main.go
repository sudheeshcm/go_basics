package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5 // item added to a channel
}

func main() {
	fooVal := make(chan int, 10) // 10 is the buffer size to avoid sync issues
	// buffers will synchronize the specified number of items for the channel
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}
	wg.Wait()
	close(fooVal)
	// value1 := <-fooVal --- (LIFO) value read from a channel
	for item := range fooVal { // items read from a channel
		fmt.Println(item)
	}
}

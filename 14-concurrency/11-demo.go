package main

import (
	"fmt"
)

func main() {

	ch := add(100, 200)
	result := <-ch //receive'ing data from the channel
	fmt.Println(result)
}

func add(x, y int) chan int {
	ch := make(chan int)
	go func() {
		result := x + y
		ch <- result //send'ing data to the channel
	}()
	return ch
}

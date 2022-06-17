package main

import (
	"fmt"
)

func main() {
	//var ch chan int = make(chan int)
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch //receive'ing data from the channel
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result //send'ing data to the channel
}

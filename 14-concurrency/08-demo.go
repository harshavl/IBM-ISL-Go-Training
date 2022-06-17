package main

import (
	"fmt"
	"sync"
)

func main() {
	//var ch chan int = make(chan int)
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch //receive'ing data from the channel
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {

	result := x + y
	ch <- result //send'ing data to the channel
	wg.Done()
}

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	data := <-ch //initiating the 'receive' process & blocking until the data becomes available
	fmt.Println(data)
}

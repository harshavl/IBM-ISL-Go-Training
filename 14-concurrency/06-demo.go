package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println("counter =", counter.count)
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}

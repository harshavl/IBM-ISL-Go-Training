package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //schedule the function execution to the scheduler
	f2()

	//DO NOT DO THIS
	time.Sleep(2 * time.Second) // blocking the execution of main() so that the scheduler can execute the scheduled goroutine
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second) //simulating a time consuming operation
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}

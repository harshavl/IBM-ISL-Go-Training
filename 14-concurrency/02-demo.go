package main

import "fmt"

func main() {
	go f1() //schedule the function execution to the scheduler
	f2()

	//DO NOT DO THIS
	fmt.Scanln() // blocking the execution of main() so that the scheduler can execute the scheduled goroutine
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

/* functions passed as arguments to other functions */
package main

import (
	"fmt"
	"time"
)

func main() {
	exec(fn)

	exec(func() {
		fmt.Println("anon fn invoked")
	})

	/*
		fmt.Println("before invocation")
		add(100, 200)
		fmt.Println("after invocation")

		fmt.Println("before invocation")
		subtract(100, 200)
		fmt.Println("after invocation")
	*/

	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)

	profileOperation(100, 200, add)
	profileOperation(100, 200, subtract)
}

func fn() {
	fmt.Println("fn invoked")
}

func exec(f func()) {
	f()
}

func profileOperation(x, y int, oper func(int, int)) {
	start := time.Now()
	oper(x, y)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Elapsed : ", elapsed)
}

func logOperation(x, y int, oper func(int, int)) {
	fmt.Println("before invocation")
	oper(x, y)
	fmt.Println("after invocation")
}

func logAdd(x, y int) {
	fmt.Println("before invocation")
	add(x, y)
	fmt.Println("after invocation")
}

func logSubtract(x, y int) {
	fmt.Println("before invocation")
	subtract(x, y)
	fmt.Println("after invocation")
}

func add(x, y int) {
	time.Sleep(3 * time.Second)
	fmt.Println("add result = ", x+y)
}

func subtract(x, y int) {
	time.Sleep(5 * time.Second)
	fmt.Println("subtract result = ", x-y)
}

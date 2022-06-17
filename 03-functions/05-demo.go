/* functions returned as return results */

package main

import (
	"fmt"
	"time"
)

func main() {
	/* fn := getFn()
	fn()

	add := getAdd()
	add(100, 200) */

	/*
		logAdd := logOperation(add)
		logAdd(100, 200)

		logSubtract := logOperation(subtract)
		logSubtract(100, 200)
	*/

	//combining log with profiling
	/*
		logAdd := logOperation(add)
		profileWithLogAdd := profileOperation(logAdd)
		profileWithLogAdd(100, 200)

		logSubtract := logOperation(subtract)
		profileWithLogSubtract := profileOperation(logSubtract)
		profileWithLogSubtract(100, 200)
	*/
	profileWithLogAdd := profileOperation(logOperation(add))
	profileWithLogAdd(100, 200)

	profileWithLogSubtract := profileOperation(logOperation(subtract))
	profileWithLogSubtract(100, 200)

}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}

func getAdd() func(int, int) {
	return func(x, y int) {
		fmt.Println("add Result = ", x+y)
	}
}

func logOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("before invocation")
		oper(x, y)
		fmt.Println("after invocation")
	}
}

func profileOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Now().Sub(start)
		fmt.Println("Elapsed =", elapsed)
	}
}

func add(x, y int) {
	time.Sleep(3 * time.Second)
	fmt.Println("add result = ", x+y)
}

func subtract(x, y int) {
	time.Sleep(5 * time.Second)
	fmt.Println("subtract result = ", x-y)
}

/* anonymous functions */

package main

import "fmt"

func main() {

	/* simple anonymous function */
	msg := "fn invoked"
	func() {
		fmt.Println(msg)
	}()

	/* with arguments */
	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	/* with return value */
	addResult := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println("addResult =", addResult)

	/* with mutiple return values */
	quotient, remainder := func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}(100, 4)
	fmt.Printf("quotient = %d, remainder = %d\n", quotient, remainder)
}

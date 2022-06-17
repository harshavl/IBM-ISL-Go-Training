/* functions assigned to variables */

package main

import "fmt"

func main() {

	//var fn func()
	fn := func() {
		fmt.Println("fn invoked")
	}
	fn()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	addResult := add(100, 200)
	fmt.Println("addResult =", addResult)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	quotient, remainder := divide(100, 4)
	fmt.Printf("quotient = %d, remainder = %d\n", quotient, remainder)
}

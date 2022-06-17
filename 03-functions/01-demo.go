package main

import "fmt"

func main() {

	fn()
	greet("Magesh")
	fmt.Print(getGreetMsg("Suresh"))
	fmt.Println(add(100, 200))

	//fmt.Println(divide(100, 7))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
	*/

	quotient, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d \n", quotient)
}

/* function with no argument & no return result */
func fn() {
	fmt.Println("fn invoked")
}

/* function with one argument */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/* function with one argument and one result */
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
}

/* function with more than one argument */
/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/* function with more than one return result */

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}

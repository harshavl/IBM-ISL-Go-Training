package main

import "fmt"

func main() {
	/*
		var msg string
		msg = "Hello World"
	*/
	/*
		var msg string = "Hello World!"
	*/

	/*
		//type inference
		var msg = "Hello World!"
	*/
	msg := "Hello World!" // := syntax is applicable ONLY in function scope (NOT in package level)
	fmt.Println(msg)

	//multiple variable declarations
	/*
		var x int
		var y int
		var result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var result int = x + y
		var str string = "Sum of 100 and 200 is"
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   int = 100, 200
			result int
			str    string = "Sum of 100 and 200 is"
		)
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result = 100, 200, 0
			str          = "Sum of 100 and 200 is"
		)
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		x, y, result := 100, 200, 0
		str := "Sum of 100 and 200 is"

		result = x + y
		fmt.Println(str, result)
	*/

	/*
		x, y, str := 100, 200, "Sum of 100 and 200 is"
		result := x + y
		fmt.Println(str, result)
	*/

	x, y := 100, 200
	result := x + y
	fmt.Printf("Sum of x(%d)[%T] and y(%d)[%T] is %d\n", x, x, y, y, result)

	//constants
	const pi float32 = 3.14

	//iota
	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota
			green
			_
			blue
		)
	*/

	/*
		const (
			red = iota + 2
			green
			blue
		)
	*/

	const (
		red = iota * 2
		green
		blue
	)

	fmt.Printf("red = %v, green = %v, blue = %v\n", red, green, blue)

	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
}

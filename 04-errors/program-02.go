package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be 0")

func main() {
	q, r, err := divide(100, 7)
	if err == DivideByZeroError {
		fmt.Println("dont attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("something went wrong!", err)
		return
	}
	fmt.Println(q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be 0")
		return 0, 0, err
	}
	quotient, remainder := x/y, x%y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient, remainder = x/y, x%y
	return
}

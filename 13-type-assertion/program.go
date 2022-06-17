package main

import "fmt"

func main() {
	//var x interface{}
	var x any

	//x = 100
	//x = "This is a string"
	//x = true
	x = float32(1000.876)
	//x = struct{}{}
	//x = []int{3, 1, 4, 2, 5}
	fmt.Println(x)

	if no, ok := x.(int); ok {
		y := no + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, 2x = ", 2*val)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is a boolean, !x = ", !val)
	case float32:
		fmt.Println("x is a float32, x - 0.01 = ", val-0.01)
	case struct{}:
		fmt.Println("x is an emptry struct")
	default:
		fmt.Println("x is an unknown type")
	}
}

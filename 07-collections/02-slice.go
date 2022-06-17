package main

import "fmt"

func main() {
	//var nos []int
	/*
		var nos []int = []int{}
		nos[0] = 100
	*/
	//var nos []int = []int{3, 1, 4, 2, 5}
	//var nos  = []int{3, 1, 4, 2, 5}
	//nos := []int{3, 1, 4, 2, 5}

	nos := make([]int, 5)
	nos[0] = 3
	nos[1] = 1
	nos[2] = 4
	nos[3] = 2
	nos[4] = 5
	fmt.Println(nos)

	fmt.Println("Iterating a slice")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println()
	fmt.Println("Iterating a slice (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	fmt.Println("After appending a value")
	nos = append(nos, 10)
	fmt.Println(nos)

	fmt.Println("After appending a multiple values")
	nos = append(nos, 20, 30, 40)
	fmt.Println(nos)

	fmt.Println("After appending another slice")
	hunderds := []int{100, 200, 300}
	nos = append(nos, hunderds...)
	fmt.Println(nos)
}

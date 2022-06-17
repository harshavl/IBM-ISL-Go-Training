package main

import "fmt"

func main() {
	//nos := make([]int, 5)
	//nos := []int{3, 1, 4, 2, 5}

	nos := make([]int, 5 /* initialized */, 20 /* memory allocated */)
	nos[0] = 3
	nos[1] = 1
	nos[2] = 4
	nos[3] = 2
	nos[4] = 5
	fmt.Printf("len(nos)=%d, cap(nos)=%d, nos=%v\n", len(nos), cap(nos), nos)

	nos = append(nos, 10, 20, 30, 40, 50)
	fmt.Printf("len(nos)=%d, cap(nos)=%d, nos=%v\n", len(nos), cap(nos), nos)

	nos = append(nos, 100)
	fmt.Printf("len(nos)=%d, cap(nos)=%d, nos=%v\n", len(nos), cap(nos), nos)

	/*
		slicing
		[lo:hi] => from lo to hi-1
		[lo:] => from lo to end of the slice
		[:hi] => from 0 to hi - 1
	*/

	fmt.Println("nos[3:6] =", nos[3:6])
	fmt.Println("nos[3:] =", nos[3:])
	fmt.Println("nos[:6] =", nos[:6])

	newNos := nos
	nos[0] = 500
	fmt.Printf("nos = %v, newNos = %v\n", nos, newNos)

	subset := nos[3:6]
	subset[0] = 2000
	fmt.Printf("nos = %v, subset = %v\n", nos, subset)
}

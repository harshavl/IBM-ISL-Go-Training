package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no
	fmt.Println(noPtr, no)

	//dereferencing - access the underlying value at an address
	var valueAtPtr = *noPtr
	fmt.Println(valueAtPtr)

	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	x, y := 100, 200
	fmt.Printf("Before swapping, x = %d and y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)

	nos := []int{3, 1, 4, 2, 5}
	fmt.Println("Before sorting, nos =", nos)
	sort(nos)
	fmt.Println("After sorting, nos =", nos)

}

func increment(x *int) {
	*x += 1
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}

func sort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

package main

import "fmt"

func main() {
	//var productRanks map[string]int
	//var productRanks map[string]int = make(map[string]int)
	//var productRanks = make(map[string]int)

	/*
		productRanks := make(map[string]int)
		productRanks["pen"] = 3
	*/
	/*
		productRanks := map[string]int{}
		productRanks["pen"] = 3
	*/

	//productRanks := map[string]int{"Pen": 3, "Marker": 2, "Pencil": 1}

	productRanks := map[string]int{
		"Pen":    3,
		"Marker": 2,
		"Pencil": 1,
		"Mouse":  4,
	}
	fmt.Println(productRanks)

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println()
	fmt.Println("Checking if a key exists")
	//keyToCheck := "Marker"
	keyToCheck := "Sharpner"
	if val, exists := productRanks[keyToCheck]; !exists {
		fmt.Printf("Key [%q] doesnot exist in %v\n", keyToCheck, productRanks)
	} else {
		fmt.Printf("value of productRanks[%q] is %d\n", keyToCheck, val)
	}

	fmt.Println()
	fmt.Println("Removing a key/value pair")
	//keyToRemove := "Pen"
	keyToRemove := "Sharper"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing key [%q], productRanks = %v\n", keyToRemove, productRanks)
}

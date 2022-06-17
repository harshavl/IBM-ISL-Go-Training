package main

import "fmt"

func main() {
	//if else
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/
	fmt.Println("Using if-else")
	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}
	fmt.Println()

	fmt.Println("Switch Case")
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/

	score := 6
	/*
		switch score {
		case 0:
			fmt.Println("Terrible")
		case 1:
			fmt.Println("Terrible")
		case 2:
			fmt.Println("Terrible")
		case 3:
			fmt.Println("Terrible")
		case 4:
			fmt.Println("Not Bad")
		case 5:
			fmt.Println("Not Bad")
		case 6:
			fmt.Println("Not Bad")
		case 7:
			fmt.Println("Not Bad")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid Score")
		}
	*/
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid Score")
	}

	x := 15
	switch {
	case x%2 == 0:
		fmt.Printf("x is an even number\n")
	case x%2 != 0:
		fmt.Printf("x is an odd number\n")
	}

	fmt.Println("Switch case with fallthrough")
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
	}

	plan := "Pro"
	switch plan {
	case "Premium":
		fmt.Println("All Premium features")
		fallthrough
	case "Super":
		fmt.Println("All Super features")
		fallthrough
	case "Pro":
		fmt.Println("All Pro features")
		fallthrough
	case "Free":
		fmt.Println("All Free features")
	default:
		fmt.Println("Invalid plan")
	}
	fmt.Println()

	fmt.Println("for constructs")
	fmt.Println("v1.0")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("v2.0 (while)")
	num := 1
	for num < 100 {
		num += num
	}
	fmt.Println("num = ", num)

	fmt.Println("v3.0 (infinite)")
	numSum := 1
	for {
		numSum += numSum
		if numSum > 100 {
			break
		}
	}
	fmt.Println("numSum = ", numSum)

	fmt.Println("Using labels")
LOOP:
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Println(i, j)
			if i == j {
				fmt.Println("-------------")
				continue LOOP
			}
		}
	}

}

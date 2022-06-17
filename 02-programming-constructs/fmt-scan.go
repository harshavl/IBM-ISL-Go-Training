package main

import "fmt"

func main() {
	var firstName, lastName string
	//fmt.Println("Enter you name  :")
	//fmt.Scanln(&firstName, &lastName)
	//fmt.Scanf("%s %s\n", &firstName, &lastName)

	fmt.Println("Enter you name [LastName,FirstName] :")
	fmt.Scanf("%s , %s\n", &lastName, &firstName)
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)

	var no int
	fmt.Println("Enter a number")
	fmt.Scanln(&no)
	fmt.Printf("Is %d an even number ? :%t\n", no, no%2 == 0)
}

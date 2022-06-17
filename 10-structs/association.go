package main

import "fmt"

type Organisation struct {
	Name string
	City string
}

type Employee struct {
	Id   int
	Name string
	City string
	Org  *Organisation
}

func main() {
	ibm := Organisation{
		Name: "IBM",
		City: "Pune",
	}
	emp1 := Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bengaluru",
		Org:  &ibm,
	}
	fmt.Printf("%#v\n", emp1)

	emp2 := Employee{
		Id:   101,
		Name: "Suresh",
		City: "Cochin",
		Org:  &ibm,
	}
	fmt.Printf("%#v\n", emp2)

	fmt.Println("Changing Org city")
	ibm.City = "Noida"
	fmt.Printf("ibm = %#v\n", ibm)
	fmt.Printf("emp1.Org = %v\n", *emp1.Org)
	fmt.Printf("emp2.Org = %v\n", *emp2.Org)
}

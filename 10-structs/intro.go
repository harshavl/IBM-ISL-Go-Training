package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func main() {
	//var emp Employee
	//fmt.Println(emp)
	//var emp Employee = Employee{100, "Magesh", "Bengaluru"}
	//var emp Employee = Employee{Id: 100, Name: "Magesh"}
	/*
		var emp Employee = Employee{
			Id:   100,
			Name: "Magesh",
		}
	*/
	//emp := Employee{} // equivalent to "var emp Employee"

	/*
		emp := Employee{
			Id:   100,
			Name: "Magesh",
		}
		fmt.Printf("%#v\n", emp)
		//accessing the attributes
		fmt.Println("Employee Name : ", emp.Name)
	*/

	//using pointers
	//emp := new(Employee)

	emp := &Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bengaluru",
	}
	fmt.Printf("%#v\n", emp)
	//fmt.Println("Employee Name : ", (*emp).Name)
	fmt.Println("Employee Name : ", emp.Name)

}

package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}

	pen.WhoAmI()

	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())

	fmt.Println("After applying 10% discount")
	//ApplyDiscount(&pen, 10)
	//(&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10)
	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())
}

func (Product) WhoAmI() {
	fmt.Println("I am a product")
}

//As a function
/*
func Format(product Product) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%v, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}
*/

//As a method
func (product Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%v, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

//As a function
/*
func ApplyDiscount(product *Product, discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}
*/

//As a method
func (product *Product) ApplyDiscount(discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

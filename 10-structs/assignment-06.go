/*
	Create a product struct with the following fields
		Id
		Name
		Cost
		Units
		Category

	Create the following utility functions for the product
		Format()
			Returns a formatted string of the given product
			ex : Id=100, Name="Pen", Cost=10, Units=100, Category="Stationary"

		ApplyDiscount(...)
			apply the given discount (%) for the given product
			ex:
				Before discount, if product cost = 10
				apply discount of 10%
				After discount, product cost should be 9

*/

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
	fmt.Println(Format(pen))

	fmt.Println("After applying 10% discount")
	ApplyDiscount(&pen, 10)
	fmt.Println(Format(pen))
}

func Format(product Product) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%v, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func ApplyDiscount(product *Product, discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

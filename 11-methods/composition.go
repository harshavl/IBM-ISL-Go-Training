package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Units:    units,
			Category: category,
		},
		Expiry: expiry,
	}
}

func main() {

	grapes := NewPerishableProduct(101, "Grapes", 50, 10, "Food", "5 Days")
	fmt.Println(grapes.Product.Name)
	fmt.Println(grapes.Name)

	//fmt.Println(FormatPerishableProduct(grapes))
	//fmt.Println(grapes.Product.Format())
	fmt.Println(grapes.Format())

	//ApplyDiscount(&grapes.Product, 10)
	fmt.Println("After applying 10% discount")
	//grapes.Product.ApplyDiscount(10)
	grapes.ApplyDiscount(10)

	//fmt.Println(FormatPerishableProduct(grapes))
	//fmt.Println(grapes.Product.Format())
	fmt.Println(grapes.Format())
}

/*
func Format(product Product) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%v, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}
*/

func (product Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%v, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

/*
func ApplyDiscount(product *Product, discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}
*/

func (product *Product) ApplyDiscount(discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

/*
func FormatPerishableProduct(pp PerishableProduct) string {
	return fmt.Sprintf("%v, Expiry=%q", pp.Product.Format(), pp.Expiry)
}
*/

//overriding the Product.Format() method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%v, Expiry=%q", pp.Product.Format(), pp.Expiry)
}

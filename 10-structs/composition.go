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
	/* 	Dummy */
	Expiry string
}

type Dummy struct {
	Cost float32
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
	/* grapes := PerishableProduct{
		Product{101, "Grapes", 50, 10, "Food"},
		"5 Days",
	} */

	/*
		grapes := PerishableProduct{
			Product: Product{101, "Grapes", 50, 10, "Food"},
			Expiry:  "5 Days",
		}
	*/
	grapes := NewPerishableProduct(101, "Grapes", 50, 10, "Food", "5 Days")
	fmt.Println(grapes)
	fmt.Println("grapes Expiry = ", grapes.Expiry)
	//fmt.Println("grapes Cost =", grapes.Product.Cost)
	fmt.Println("grapes Cost =", grapes.Cost)

	fmt.Println(FormatPerishableProduct(grapes))
	ApplyDiscount(&grapes.Product, 10)
	fmt.Println(FormatPerishableProduct(grapes))
}

/* Do NOT modify the below functions */
func Format(product Product) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%v, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func ApplyDiscount(product *Product, discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

func FormatPerishableProduct(pp PerishableProduct) string {
	return fmt.Sprintf("%v, Expiry=%q", Format(pp.Product), pp.Expiry)
}

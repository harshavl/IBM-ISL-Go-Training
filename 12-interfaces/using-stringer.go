package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/*
func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}
*/

//implementing fmt.Stringer interface
func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
	Write the apis for the following

	IndexOf => return the index of the given product
		ex:  returns the index of the given product

	Includes => return true if the given product is present in the collection else return false
		ex:  returns true if the given product is present in the collection

	Filter => return a new collection of products that satisfy the given condition
		use cases:
			1. filter all costly products (cost > 1000)
				OR
			2. filter all stationary products (category = "Stationary")
				OR
			3. filter all costly stationary products
			etc

	Any => return true if any of the product in the collections satifies the given criteria
		use cases:
			1. are there any costly product (cost > 1000)?
			2. are there any stationary product (category = "Stationary")?
			3. are there any costly stationary product?
			etc

	All => return true if all the products in the collections satifies the given criteria
		use cases:
			1. are all the products costly products (cost > 1000)?
			2. are all the products stationary products (category = "Stationary")?
			3. are all the products costly stationary products?
			etc

*/

type Products []Product

/*
func (products Products) Format() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%v\n", product.Format())
	}
	return result
}
*/
//implementing fmt.Stringer interface
func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%v\n", product)
	}
	return result
}

func (products Products) IndexOf(product Product) int {
	for idx, p := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	for _, p := range products {
		if p == product {
			return true
		}
	}
	return false
}

//Day-1
func (products Products) Filter(criteria func(Product) bool) Products {
	result := Products{}
	for _, product := range products {
		if criteria(product) {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) All(criteria func(Product) bool) bool {
	for _, product := range products {
		if !criteria(product) {
			return false
		}
	}
	return true
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

//utility functions

/*
func Or(leftPredicate func(product Product) bool, rightPredicate func(product Product) bool) func(product Product) bool {
	return func(product Product) bool {
		return leftPredicate(product) || rightPredicate(product)
	}
}

func And(leftPredicate func(product Product) bool, rightPredicate func(product Product) bool) func(product Product) bool {
	return func(product Product) bool {
		return leftPredicate(product) && rightPredicate(product)
	}
}
*/

type ProductPredicate func(product Product) bool

func Or(leftPredicate ProductPredicate, rightPredicate ProductPredicate) ProductPredicate {
	return func(product Product) bool {
		return leftPredicate(product) || rightPredicate(product)
	}
}

func And(leftPredicate ProductPredicate, rightPredicate ProductPredicate) ProductPredicate {
	return func(product Product) bool {
		return leftPredicate(product) && rightPredicate(product)
	}
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	//fmt.Println(products.Format())
	fmt.Println(products)

	stove := Product{102, "Stove", 5000, 5, "Utencil"}
	fmt.Println("IndexOf stove = ", products.IndexOf(stove))
	fmt.Println("Does products include stove ?:", products.Includes(stove))

	fmt.Println()
	fmt.Println("Filter")
	fmt.Println("Costly Products [cost > 1000]")
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	//fmt.Println(costlyProducts.Format())
	fmt.Println(costlyProducts)

	fmt.Println("Stationary products [category = Stationary]")
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	//fmt.Println(stationaryProducts.Format())
	fmt.Println(stationaryProducts)

	fmt.Println("Costly Stationary products")
	costlyStationaryProductPredicate := And(costlyProductPredicate, stationaryProductPredicate)
	costlyStationaryProducts := products.Filter(costlyStationaryProductPredicate)
	//fmt.Println(costlyStationaryProducts.Format())
	fmt.Println(costlyStationaryProducts)

	fmt.Println("All")
	fmt.Println("Are all products costly products ?:", products.All(costlyProductPredicate))

	fmt.Println("Any")
	fmt.Println("Are there any costly products ?:", products.Any(costlyProductPredicate))

	//Day-10
	utencilProducts := products.Filter(func(product Product) bool {
		return product.Category == "Utencil"
	})
	//fmt.Println(utencilProducts.Format())
	fmt.Println(utencilProducts)
}

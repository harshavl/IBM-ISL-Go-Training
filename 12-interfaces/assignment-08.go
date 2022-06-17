package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

//implementing fmt.Stringer interface
func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
	Write the apis for the following
		Sort => Sort the products collection by any attribute
			IMPORTANT : MUST Use sort.Sort() function
            use cases:
                1. sort the products collection by cost
                2. sort the products collection by name
                3. sort the products collection by units
                4. sort the products collection by cost in descending order
                5. sort the products collection by name in descending order
                6. sort the products collection by units in descending order
				etc

*/

type Products []Product

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

//implementation of the 'Interface' interface
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

//by name
type ByName struct {
	Products
}

//overriding the "Less()" method of Products
func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//by cost
type ByCost struct {
	Products
}

//overriding the "Less()" method of Products
func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

//by units
type ByUnits struct {
	Products
}

//overriding the "Less()" method of Products
func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

//by category
type ByCategory struct {
	Products
}

//overriding the "Less()" method of Products
func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Category":
		sort.Sort(ByCategory{products})
	}
}

func (products Products) SortSlice(attrName string) {
	switch attrName {
	case "Id":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Id < products[j].Id
		})
	case "Name":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
	case "Cost":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Cost < products[j].Cost
		})
	case "Units":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Units < products[j].Units
		})
	case "Category":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Category < products[j].Category
		})
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

	/*
		fmt.Println()
		fmt.Println("Sorting [default - by id]")
		sort.Sort(products)
		fmt.Println(products)

		fmt.Println()
		fmt.Println("Sorting [by name]")
		sort.Sort(ByName{products})
		fmt.Println(products)

		fmt.Println()
		fmt.Println("Sorting [by cost]")
		sort.Sort(ByCost{products})
		fmt.Println(products)

		fmt.Println()
		fmt.Println("Sorting [by units]")
		sort.Sort(ByUnits{products})
		fmt.Println(products)

		fmt.Println()
		fmt.Println("Sorting [by category]")
		sort.Sort(ByCategory{products})
		fmt.Println(products)
	*/

	/* //using products.Sort()
	fmt.Println()
	fmt.Println("Sorting [default - by id]")
	products.Sort("Id")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting [by name]")
	products.Sort("Name")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting [by cost]")
	products.Sort("Cost")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting [by units]")
	products.Sort("Units")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting [by category]")
	products.Sort("Category")
	fmt.Println(products)
	*/

	//using products.SortSlice()
	fmt.Println()
	fmt.Println("Sorting [default - by id]")
	products.SortSlice("Id")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting [by name]")
	products.SortSlice("Name")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting [by cost]")
	products.SortSlice("Cost")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting [by units]")
	products.SortSlice("Units")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting [by category]")
	products.SortSlice("Category")
	fmt.Println(products)
}

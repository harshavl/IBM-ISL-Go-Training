package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

type ShapeWithArea interface {
	Area() float32
}

func printArea(x ShapeWithArea) {
	fmt.Println("Area =", x.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func printPerimeter(x ShapeWithPerimeter) {
	fmt.Println("Perimeter =", x.Perimeter())
}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func printShape(x Shape) {
	printArea(x)
	printPerimeter(x)
}

func main() {
	c := Circle{12}
	//fmt.Println("Area =", c.Area())
	/*
		printArea(c)
		//fmt.Println("Perimeter =", c.Perimeter())
		printPerimeter(c)
	*/
	printShape(c)

	r := Rectangle{10, 12}
	//fmt.Println("Area =", r.Area())
	/*
		printArea(r)
		//fmt.Println("Perimeter =", r.Perimeter())
		printPerimeter(r)
	*/
	printShape(r)
}

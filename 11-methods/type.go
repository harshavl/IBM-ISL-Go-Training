package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {

	var i int = 100
	var f float32
	f = float32(i)
	fmt.Println(f)

	str := MyStr("This is a sample string")
	fmt.Println(str.Length())
}

package main

import "fmt"

var db string = "postgres"

// var db = "postgres"
// db := "postgres"   // error

func main() {
	var i int
	fmt.Printf("i:%v\n", i)
	i = 10
	fmt.Printf("i:%v\n", i)

	var s string
	fmt.Printf("s:%v\n", s)
	s = "hello world"
	fmt.Printf("s:%v\n", s)

	var b bool
	fmt.Printf("b:%v\n", b)
	b = true
	fmt.Printf("b:%v\n", b)

	println("==============")

	age := 20
	fmt.Printf("age:%v\n", age)

	name := "anousone"
	fmt.Printf("name:%v\n", name)
}

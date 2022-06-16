package main

import "fmt"

type Customer struct {
	name string
	age  int
}

func (c Customer) Hello() string {
	return "hello " + c.name
}

func (c *Customer) SetName() {
	c.name = "new"
}

func main() {
	c := Customer{name: "anousone", age: 23}
	fmt.Println(c)
	fmt.Println(c.name)
	c.SetName()
	fmt.Println(c.name)
	fmt.Println(c.Hello())
}

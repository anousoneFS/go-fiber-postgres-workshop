package main

import "fmt"

func main() {
	fmt.Printf("get(): %v\n", get())
	fmt.Printf("hello(): %v\n", hello("anousone"))
	fmt.Printf("add():%v\n", add(10, 12))
	number := []int{1, 2, 3, 4}
	fmt.Printf("sum():%v\n", sum(number))
	fmt.Printf("sum2(): %v\n", sum2(1, 2, 3, 4))
	fmt.Printf("sum2(): %v\n", sum2(number...))
	fmt.Println(swap("aaa", "bbb"))

	score := 60
	func(s int) {
		if s > 50 {
			fmt.Println("pass")
		} else {
			fmt.Println("not pass")
		}
	}(score)

	total := func(i ...int) int {
		t := 0
		for _, i := range i {
			t += i
		}
		return t
	}(append(number, 3, 4)...)
	fmt.Printf("total: %v\n", total)
}

func get() string {
	return "hello world"
}

func hello(name string) string {
	return "hello " + name
}

func add(a, b int) int {
	return a + b
}

func sum(a []int) int {
	t := 0
	for _, i := range a {
		t += i
	}
	return t
}

func sum2(a ...int) int {
	t := 0
	for _, i := range a {
		t += i
	}
	return t
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

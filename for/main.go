package main

import "fmt"

func main() {
	number := []int{10, 20, 30, 40, 50, 60, 70, 80}

	println("1 for ----------")
	for i := 0; i < len(number); i++ {
		fmt.Printf("index: %v, item: %v\n", i, number[i])
	}

	println("2 for each ---------")
	for index, item := range number {
		fmt.Printf("index: %v, item: %v\n", index, item)
	}

	println("3 while ---------")
	i := 0
	for {
		fmt.Printf("index: %v, item: %v\n", i, number[i])
		if i == len(number)-1 {
			break
		}
		i += 1
	}

	println("4 do while ---------")
	i = 0
	for i < len(number) {
		fmt.Printf("index: %v, item: %v\n", i, number[i])
		i += 1
	}
}

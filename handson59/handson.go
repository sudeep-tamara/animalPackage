package main

import "fmt"

func foo(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func main() {
	numSlice := []int{1, 34, 4, 4, 43, 62, 27}
	fmt.Println(foo(numSlice...))
}

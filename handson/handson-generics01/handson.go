package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func addInt(a int, b int) int {
	return a + b
}

func addF(a float64, b float64) float64 {
	return a + b
}

type MyNumber interface {
	constraints.Integer | constraints.Float
}

func addT[T MyNumber](a, b T) T {
	return a + b
}

type myAlias int

func main() {

	var n myAlias = 53
	fmt.Println(addInt(1, 2))
	fmt.Println(addF(1.2, 2.3))

	fmt.Println(addT(n, 24))
	fmt.Println(addT(1.54, 24))
	fmt.Println(addT(1.54, 24.34))
}

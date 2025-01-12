package main

import (
	"fmt"
	"math"
)

func add(a, b int64) int64 {
	return a + b
}

func subtract(a, b int64) int64 {
	return int64(math.Abs(float64(a - b)))
}

func doOp(a int64, b int64, f func(int64, int64) int64) int64 {
	return f(a, b)
}

func main() {
	a, b := int64(10), int64(20)
	sum := doOp(a, b, add)
	diff := doOp(a, b, subtract)

	mul := func(a, b int64) int64 {
		return a * b
	}

	fmt.Printf("Sum is %d, diff is %d, multiplication s %d\n", sum, diff, mul)
}

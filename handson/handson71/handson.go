package main

import "fmt"

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, a int) string {
	return fmt.Sprintf("Function called, with result, %d", f(a))
}

func main() {
	result := printSquare(square, 10)

	fmt.Println(result)
}

package main

import "fmt"

func main() {
	x()
}

var x = func() {
	fmt.Println(42)
}

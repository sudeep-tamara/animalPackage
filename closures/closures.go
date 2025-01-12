package main

import "fmt"

func main() {

	f1 := incrementor()
	incrN := f1()
	fmt.Println(incrN)

	f2 := decrementor(incrN)
	decrN := f2()
	fmt.Printf("%T, %#v\n", f2, f2())
	fmt.Printf("%T, %#v\n", decrN, decrN)

}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		fmt.Println(x)
		return x
	}
}

func decrementor(n int) func() int {
	return func() int {
		n--
		return n
	}
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("bar:", i)
	}
}

func main() {
	fmt.Println("OS:\t", runtime.GOOS)
	fmt.Println("ARch:\t", runtime.GOARCH)
	fmt.Println("Cpus:\t", runtime.NumCPU())
	fmt.Println("rtns:\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()

	fmt.Println("rtns:\t", runtime.NumGoroutine())

	bar()
	wg.Wait()

}

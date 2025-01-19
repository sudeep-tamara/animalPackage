/*
Hands-on exercise #1
● in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
code: https://github.com/GoesToEleven/go-programming
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func go1() {
	fmt.Println("Go routine1")
	wg.Done()

}

func go2() {
	fmt.Println("Go routine2")
	wg.Done()

}

func main() {

	wg.Add(2)
	go go1()
	go go2()
	wg.Wait()
}

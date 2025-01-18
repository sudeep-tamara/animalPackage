package main

import "fmt"

func main() {
	fmt.Println(paradise("Dapoli"))

	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
}

func paradise(s string) string {
	return fmt.Sprintf("My idea of paradise is %s", s)
}

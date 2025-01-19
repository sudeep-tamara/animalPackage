package main

import "fmt"

func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total = total + v
	}
	return
}
func main() {
	nums := []int{1, 2, 5, 3, 4}
	fmt.Println(sum(nums))
}

/*
Hands-on exercise #4
Starting with this code, sort the []int and []string for each person.
solution: https://play.golang.org/p/jz
_
llY1aPp
*/
package main

import (
	"fmt"
	"sort"
)

//type xSlice []int64
//
//
//
//func (n xSlice) Len() int           { return len(n) }
//func (n xSlice) Swap(i, j int)      { n[j], n[i] = n[i], n[j] }
//func (n xSlice) Less(i, j int) bool { return n[i] < n[j] }

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}

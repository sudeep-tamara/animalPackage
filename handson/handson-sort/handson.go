package main

import (
	"fmt"
	"sort"
)

type ByName []Person

type Person struct {
	Name string
	Age  int
}

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[j], n[i] = n[i], n[j] }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }

func main() {
	p1 := Person{"Evah", 8}
	p2 := Person{"Pranali", 39}
	p3 := Person{"Sudeep", 40}
	p4 := Person{"Tom", 22}
	p5 := Person{"Sara", 22}

	people := []Person{p1, p2, p3, p4, p5}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}

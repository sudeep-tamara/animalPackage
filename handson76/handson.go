package main

import "fmt"

type person struct {
	first string
}

func changeNameVal(p person, name string) person {
	p.first = name
	return p
}

func changeNamePtr(p *person, name string) {
	p.first = name
}

func main() {
	p1 := person{
		"Tony",
	}

	fmt.Println(changeNameVal(p1, "Timmy"))

	changeNamePtr(&p1, "Jack")

	fmt.Println(p1)

}

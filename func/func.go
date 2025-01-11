package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Printf("%s(type:%T) speaks\n", p.first, p)
}

func (sa secretAgent) speak() {
	fmt.Printf("Secret agent %s(type:%T) speaks\n", sa.first, sa)
}

func saySomething(h human) {
	h.speak()
}

type human interface {
	speak()
}

func main() {

	p1 := person{
		"James",
	}

	p2 := person{
		"Peter",
	}

	saySomething(p1)
	saySomething(p2)

	sa := secretAgent{
		person{
			first: "Bond",
		},
		false,
	}

	saySomething(sa)
}

package main

import "fmt"

func main() {

	p := Person{
		"Sudeep",
		23,
	}
	p.printMe()
	p.incrAge()
	p.printMe()

}

type Person struct {
	Name string
	Age  int64
}

func (p *Person) printMe() {
	fmt.Printf("Hello, %s, your age is %d\n", p.Name, p.Age)
}

func (p *Person) incrAge() {
	p.Age = p.Age + 1
}

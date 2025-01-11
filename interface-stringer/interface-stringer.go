package main

import (
	"fmt"
	"log"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("This is a book")
}

type count int

func (c count) String() string {
	return fmt.Sprintf("This is a number %d", int(c))
}

func logInfo(s fmt.Stringer) {
	log.Println("This the prefix: ", fmt.Sprint(s))
}

func main() {

	b := book{
		"Fables",
	}

	c := count(10)

	log.Println(c)
	log.Println(b)

	logInfo(c)
	logInfo(b)

	var n count = 20
	logInfo(n)
}

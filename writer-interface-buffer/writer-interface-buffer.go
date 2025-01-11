package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) WriteOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	//bb := bytes.NewBuffer([]byte("Hello World"))
	bbStr := bytes.NewBufferString("Hello ")

	fmt.Println(bbStr.String())

	bbStr.Write([]byte("World. "))
	bbStr.WriteString("New World")
	fmt.Println(bbStr.String())

	bbStr.Reset()

	bbStr.Write([]byte("World"))
	fmt.Println(bbStr.String())

	p1 := person{
		"James",
	}

	err = p1.WriteOut(bbStr)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	err = p1.WriteOut(f)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(bbStr.String())

}

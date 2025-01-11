package main

import (
	"fmt"
	"log"
	"os"
)

func errLog(err error) {
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

func main() {

	f, err := os.Create("output.txt")
	errLog(err)
	defer func() {
		err := f.Close()
		errLog(err)
	}()

	s := []byte("Hello Gophers!")

	n, err := f.Write(s)
	errLog(err)

	fmt.Printf("Wrote %d bytes to file", n)
}

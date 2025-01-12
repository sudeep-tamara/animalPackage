package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("output.txt")
	if err != nil {
		log.Fatalf("")
	}

	fmt.Printf("Data: %s, Type: %T\n", xb, xb)
	fmt.Println(string(xb))
}

func readFile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("")
	}
	return xb, nil
}

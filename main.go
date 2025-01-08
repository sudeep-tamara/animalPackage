package main

import (
	"fmt"
	"github.com/sudeep-tamara/puppy"
	"math/rand"
)

var x = "asdasd"

const y = "sadasd"

type ByteSize int64

const (
	c1 = iota
	c2
	c3
	c4
	c5
	c6
	c7
	c8
)
const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func init() {
	fmt.Println("Begin init()" + string(rune(c2)))
	fmt.Println(c3)
	fmt.Println(1 << c3)
	fmt.Println(1 << c8)
}

func main() {

	fmt.Println("Hello Gophers Again!! :)")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())
	x := rand.Intn(20)
	y := x + 10
	switch x {
	case 10:
		fmt.Printf("is 10 x:%d y:%d", x, y)
		break
	case 20:
		fmt.Printf("is 20 x:%d y:%d", x, y)
		break
	default:
		fmt.Printf("is not 10 or 20 x:%d y:%d", x, y)
	}
}

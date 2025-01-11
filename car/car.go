package main

import "fmt"

// type Engine
type Engine struct {
	typeOfEngine rune
	cylinder     int
	specs        map[string]int
}

// type Car
type Car struct {
	brand string
	Engine
	parts      [4]string
	extraParts []string
}

func (e *Engine) Start() {
	fmt.Println("Start car")
}

func (e Engine) Stop() {
	fmt.Println("Stop Car")
}

//function on engine to start

func main() {
	engineSpecs1 := make(map[string]int)
	engineSpecs1["hp"] = 100
	engineSpecs1["torque"] = 150

	engineSpecs2 := map[string]int{
		"hp":     250,
		"torque": 170,
	}

	e1 := Engine{
		'v',
		12,
		engineSpecs1,
	}

	e2 := Engine{
		'v',
		8,
		engineSpecs2,
	}

	c1 := Car{
		"TATA",
		e1,
		[4]string{"windows", "gears", "wheels", "seats"},
		[]string{"mirrors", "brake"},
	}

	c2 := Car{
		"Kia",
		e2,
		[4]string{"part1", "part2", "part3", "part4"},
		[]string{},
	}

	c3 := Car{
		"Toyota",
		Engine{
			'v',
			22,
			map[string]int{
				"hp":     20000,
				"torque": 3000,
			},
		},
		[4]string{"part1", "part2", "part3", "part4"},
		[]string{},
	}

	c1.Start()
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
}

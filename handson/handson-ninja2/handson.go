package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	var people []person
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		//fmt.Println(err)
	}
	//fmt.Println(people)

	for _, person := range people {
		fmt.Printf("\n====\nName: %s %s, Age: %d\n", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Printf("\t%s\n", saying)
		}
		fmt.Printf("====")
	}
}

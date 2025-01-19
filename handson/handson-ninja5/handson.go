/*
Hands-on exercise #5
Starting with this code, sort the []user by
● age
● last

Also sort each []string “Sayings” for each user
● print everything in a way that is pleasant
solution: https://play.golang.org/p/8RKkdtLl6w
*/
package main

import (
	"fmt"
	"sort"
)

type ByLast []user
type ByAge []user

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (n ByLast) Len() int           { return len(n) }
func (n ByLast) Swap(i, j int)      { n[j], n[i] = n[i], n[j] }
func (n ByLast) Less(i, j int) bool { return n[i].Last < n[j].Last }

func (n ByAge) Len() int           { return len(n) }
func (n ByAge) Swap(i, j int)      { n[j], n[i] = n[i], n[j] }
func (n ByAge) Less(i, j int) bool { return n[i].Age < n[j].Age }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	printUsers(users)

	fmt.Println("\n=========== Sorted by Name ===========\n")
	sort.Sort(ByLast(users))
	printUsers(users)

	fmt.Println("\n=========== Sorted by Age ===========\n")
	sort.Sort(ByAge(users))
	printUsers(users)

}

func printUsers(users []user) {
	for _, u := range users {
		fmt.Printf("\n Name: %s %s, \n Age : %d\nSayings:\n", u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for i, saying := range u.Sayings {
			fmt.Printf("\t %d) %s\n", i, saying)
		}
	}
}

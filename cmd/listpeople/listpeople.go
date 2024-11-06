package main

import (
	"flag"
	"fmt"
	"github.com/rickb777/go-tutorial1/fixtures"
)

var (
	number int
)

func main() {
	flag.IntVar(&number, "n", 0, "number of people")
	flag.Parse()

	fmt.Println("Adults")
	for i := 0; i < number; i++ {
		person := fixtures.RandomAdult(fixtures.RandomMaleOrFemale())
		if len(person.Addresses) > 0 {
			fmt.Printf("%3d: %s\n     %s (since %s)\n", i, person, person.Addresses[0], person.MovedIn[0])
		} else {
			fmt.Printf("%3d: %s\n     %s (since %s)\n", i, person, person.Addresses[0], person.MovedIn[0])
		}
	}

	fmt.Println("Children")
	for i := 0; i < number; i++ {
		person := fixtures.RandomChild(fixtures.RandomMaleOrFemale())
		fmt.Printf("%3d: %s\n", i, person)
	}

}

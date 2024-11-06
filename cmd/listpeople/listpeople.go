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

	for i := 0; i < number; i++ {
		person := fixtures.RandomAdult(fixtures.RandomMaleOrFemale())
		fmt.Printf("%3d: %s\n     %s\n", i, person, person.Addresses[0])
	}

}

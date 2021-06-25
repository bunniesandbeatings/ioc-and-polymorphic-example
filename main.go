package main

import (
	"fmt"

	"example.com/patterns/api"
	"example.com/patterns/processor"
)


// Behaviour


// Encapsulation



func main() {
	objects := []interface{}{
		&api.Place{
			Name: "Glastonbury",
		},
		&api.Person{
			Name:       "Jack",
			Occupation: "Footballer",
			Place:      "Totenham",
		},
		&api.Person{
			Name:       "Mary",
			Occupation: "Cartographer",
			Place:      "Glastonbury",
		},
		&api.Place{
			Name: "Totenham",
		},
	}

	processableObjects := processor.MakeProcessable(objects)

	processableObjects.Process()

	fmt.Println("\n ###########################")
	fmt.Println("# Proving it")
	fmt.Println("###########################\n")

	processableObjects.Process()
}

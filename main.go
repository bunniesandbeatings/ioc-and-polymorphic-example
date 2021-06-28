package main

import (
	"fmt"

	"example.com/patterns/api"
	"example.com/patterns/processor"
)

func main() {
	// Plain old data
	// Not immutable, but doesn't matter if this action is event driven
	// eg: state in -> state out by the end of this function
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

	fmt.Println("\n###########################")
	fmt.Println("# Data Objects")
	fmt.Println("###########################\n")
	printPodos(objects)

	processableObjects := processor.MakeProcessable(objects)

	processableObjects.Process()

	fmt.Println("\n###########################")
	fmt.Println("# Proving it")
	fmt.Println("###########################\n")

	processableObjects.Process()

	fmt.Println("\n###########################")
	fmt.Println("# mutated but not decorated")
	fmt.Println("###########################\n")
	printPodos(objects)

}

func printPodos(objects []interface{}) {
	for _, podo := range objects {
		switch v := podo.(type) {
		case *api.Place:
			fmt.Println("Place:")
			fmt.Printf("\tName:%s\n", v.Name)
			fmt.Printf("\tLocation:%+v\n", v.Location)
		case *api.Person:
			fmt.Println("Person:")
			fmt.Printf("\tName:%s\n", v.Name)
			fmt.Printf("\tOccupation:%s\n", v.Occupation)
			fmt.Printf("\tPlace:%s\n", v.Place)
			fmt.Printf("\tLocation:%+v\n", v.Location)
		}
	}
}

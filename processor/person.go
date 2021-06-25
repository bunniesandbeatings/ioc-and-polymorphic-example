package processor

import (
	"fmt"

	"example.com/patterns/api"
)

type processablePerson struct {
	dirty  bool
	object *api.Person
}

func NewPersonNode(person *api.Person) ProcessableItem {
	return &processablePerson{
		object: person,
	}
}

func (p *processablePerson) Process(sharedContext *Context) bool {
	fmt.Printf("  Person: %s\n", p.object.Name)
	if p.object.Location != nil {
		fmt.Printf("    Location: %s\n", p.object.Location)
		return false
	}

	p.object.Location = sharedContext.ResolveLocation(p.object.Place)

	// still dirty?
	return p.object.Location == nil
}

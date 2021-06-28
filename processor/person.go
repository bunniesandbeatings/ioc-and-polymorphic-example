package processor

import (
	"fmt"

	"example.com/patterns/api"
)

// Avoid metadata in the pure-data, just refer to it

// This type supports maintains state that pertains
// to ONLY the data type it wraps, eg dirty
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

	// State transition to 'ready to commit' would have happened here
	// eg:
	if p.object.Location == nil {
		return true
	}

	// persistance.Commit(p.Object)

	// still dirty?
	return false
}

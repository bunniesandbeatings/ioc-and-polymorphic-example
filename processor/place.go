package processor

import (
	"fmt"

	"example.com/patterns/api"
	"example.com/patterns/place"
)

// Avoid metadata in the pure-data, just refer to it

type processablePlace struct {
	dirty  bool
	object *api.Place
}

func NewPlaceNode(place *api.Place) ProcessableItem {
	return &processablePlace{
		object: place,
	}
}

func (p *processablePlace) Process(sharedContext *Context) bool {
	fmt.Printf("  Place: %s\n", p.object.Name)

	if place.Valid(p.object) {
		return false
	}

	place.Update(p.object)
	sharedContext.AddLocationByPlace(p.object)

	return true
}

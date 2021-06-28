package place

import (
	"fmt"
	"math/rand"

	"example.com/patterns/api"
	"example.com/patterns/processor"
)

// Avoid metadata in the pure-data, just refer to it

type processablePlace struct {
	dirty  bool
	object *api.Place
}

func NewPlaceNode(place *api.Place) processor.ProcessableItem {
	return &processablePlace{
		object: place,
	}
}

func (p *processablePlace) Process(sharedContext *processor.Context) bool {
	fmt.Printf("  Place: %s\n", p.object.Name)
	if p.object.Location != nil {
		fmt.Printf("    Location: %v\n", p.object.Location)
		return false
	}

	// This is a stochastic simulation of information becoming available on an endpoint.
	if rand.Float32() > .9 {
		p.object.Location = &api.Location{
			Lat: rand.Float64() * 100.0,
			Lon: rand.Float64() * 100.0,
		}

		sharedContext.AddLocationByPlace(p.object)
		return false
	}

	return true
}

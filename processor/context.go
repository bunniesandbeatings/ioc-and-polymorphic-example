package processor

import "example.com/patterns/api"

// Context
// The processor is not generic because Go makes that hard
// But it doesn't matter, the processor knows three things, annotated as KNOWS: in the code
type Context struct {
	// KNOWS: 1. what optimization is needed to reduce nested loops
	locations map[string]*api.Location
}

func NewContext() *Context {
	return &Context{
		locations: make(map[string]*api.Location),
	}
}

func (c *Context) ResolveLocation(placeName string) *api.Location {
	if location, ok := c.locations[placeName]; ok {
		return location
	}
	return nil
}

func (c *Context) AddLocationByPlace(place *api.Place) {
	c.locations[place.Name] = place.Location
}

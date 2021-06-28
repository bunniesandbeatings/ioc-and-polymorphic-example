package place

import (
	"math/rand"

	"example.com/patterns/api"
)

func Valid(place *api.Place) bool {
	return place.Location != nil
}


func Update(place *api.Place)  {
	if rand.Float32() > .9 {
		place.Location = &api.Location{
			Lat: rand.Float64() * 100.0,
			Lon: rand.Float64() * 100.0,
		}
	}
}

package person

import "example.com/patterns/api"

func Valid(person *api.Person) bool {
	return person.Location != nil
}
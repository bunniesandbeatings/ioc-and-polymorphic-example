package processor

import (
	"fmt"

	"example.com/patterns/api"
	"example.com/patterns/processor/person"
	"example.com/patterns/processor/place"
)

type ProcessableItem interface {
	Process(sharedContext *Context) bool
}

type ProcessableList interface {
	Process()
}

type processableList struct {
	items   []ProcessableItem
	context *Context
}

// KNOWS: 3. that we are collapsing incomplete state.
func (l *processableList) Process() {
	l.context = NewContext()

	for {
		dirty := false
		fmt.Println("--- Reconciling")
		for _, item := range l.items {
			dirtyObject := item.Process(l.context)
			if dirtyObject {
				dirty = true
			}
		}
		if !dirty {
			break
		}
	}
	fmt.Println("--- Reconciled")
}

// MakeProcessable
// KNOWS: 2. How to build node types
// This could be abstracted away as well but feels uneccessary
func MakeProcessable(items []interface{}) ProcessableList {
	var processableItems []ProcessableItem
	var processableItem ProcessableItem

	for _, item := range items {
		switch v := item.(type) {
		case *api.Place:
			processableItem = place.NewPlaceNode(v)
		case *api.Person:
			processableItem = person.NewPersonNode(v)
		default:
			panic("Provided unknown processable")
		}
		processableItems = append(processableItems, processableItem)
	}
	return &processableList{items: processableItems}
}

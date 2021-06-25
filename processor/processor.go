package processor

import (
	"fmt"

	"example.com/patterns/api"
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

func MakeProcessable(items []interface{}) ProcessableList {
	var processableItems []ProcessableItem
	var processableItem ProcessableItem

	for _, item := range items {
		switch v := item.(type) {
		case *api.Place:
			processableItem = NewPlaceNode(v)
		case *api.Person:
			processableItem = NewPersonNode(v)
		default:
			panic("Provided unknown processable")
		}
		processableItems = append(processableItems, processableItem)
	}
	return &processableList{items: processableItems}
}

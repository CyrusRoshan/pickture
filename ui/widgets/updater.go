package widgets

import (
	"reflect"

	"github.com/gotk3/gotk3/gtk"
)

// A simple explanation of why this package exists:
//
// The idea behind this project's structure is we have state, we have logic, and
// we have the way they appear graphically, which allows users to take action.
//
// In psuedocode terms:
// apply: state -> logic(action, state)
// update: appearance -> render(state)
//
// The updater widget allows us to bind() the state to our UI's appearance.
// We do this by attaching a listener (a channel) to the end of the logic()
// function. Now, we can update the UI on every state update.
//
// Note that this won't fully rebuild the actual underlying UI library's objects
// (like React will). Instead, we rely on the per-widget defined
// implementation of the update function to provide a partial re-render.

type UpdaterWidget struct {
	InnerWidget    *gtk.Widget
	updateFunction func(newState interface{})
	lastState      interface{}
}

func NewUpdaterWidget(gtkWidget *gtk.Widget, updateFunc func(newState interface{})) UpdaterWidget {
	return UpdaterWidget{
		InnerWidget:    gtkWidget,
		updateFunction: updateFunc,
	}
}

// Update returns whether a widget was updated or not
func (u *UpdaterWidget) Update(state interface{}) bool {
	if reflect.DeepEqual(u.lastState, state) {
		return false
	}
	u.updateFunction(state)
	u.lastState = state
	return true
}

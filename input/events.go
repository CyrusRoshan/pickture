package input

import (
	"github.com/gotk3/gotk3/gdk"
)

type InputEvent string

const UndoEvent = InputEvent("undo")
const NextEvent = InputEvent("next")
const QPressEvent = InputEvent("press_q")
const WPressEvent = InputEvent("press_w")
const EPressEvent = InputEvent("press_e")
const APressEvent = InputEvent("press_a")
const SPressEvent = InputEvent("press_s")
const DPressEvent = InputEvent("press_d")

var inputEvents = make(chan InputEvent, 3)

type keyEvent struct {
	Key   uint
	Event InputEvent
}

var letterKeyEvents = []keyEvent{
	keyEvent{
		Key:   gdk.KEY_Q,
		Event: QPressEvent,
	},
	keyEvent{
		Key:   gdk.KEY_W,
		Event: WPressEvent,
	},
	keyEvent{
		Key:   gdk.KEY_E,
		Event: EPressEvent,
	},
	keyEvent{
		Key:   gdk.KEY_A,
		Event: APressEvent,
	},
	keyEvent{
		Key:   gdk.KEY_S,
		Event: SPressEvent,
	},
	keyEvent{
		Key:   gdk.KEY_D,
		Event: DPressEvent,
	},
}

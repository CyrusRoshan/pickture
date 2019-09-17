package input

import (
	"github.com/faiface/pixel/pixelgl"
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

var eventChannel = make(chan InputEvent, 3)

func GetKeyPressEvents() chan InputEvent {
	return eventChannel
}

type keyButtonEventCombo struct {
	Button pixelgl.Button
	Name   string
	Event  InputEvent
}

var keyCombos = []keyButtonEventCombo{
	keyButtonEventCombo{
		Button: pixelgl.KeyQ,
		Name:   "Q",
		Event:  QPressEvent,
	},
	keyButtonEventCombo{
		Button: pixelgl.KeyW,
		Name:   "W",
		Event:  WPressEvent,
	},
	keyButtonEventCombo{
		Button: pixelgl.KeyE,
		Name:   "E",
		Event:  EPressEvent,
	},
	keyButtonEventCombo{
		Button: pixelgl.KeyA,
		Name:   "A",
		Event:  APressEvent,
	},
	keyButtonEventCombo{
		Button: pixelgl.KeyS,
		Name:   "S",
		Event:  SPressEvent,
	},
	keyButtonEventCombo{
		Button: pixelgl.KeyD,
		Name:   "D",
		Event:  DPressEvent,
	},
}

// Used to not fire multiple times in one go
var keyHasFired = map[string]bool{}
var keyTypeLastCycle = map[string]bool{}

func CalculateKeyPressChanges(win *pixelgl.Window) {
	var keyWasPressedThisCycle = false

	// Undo event
	if win.Pressed(pixelgl.KeyZ) &&
		(win.Pressed(pixelgl.KeyLeftSuper) ||
			win.Pressed(pixelgl.KeyRightSuper) ||
			win.Pressed(pixelgl.KeyLeftControl) ||
			win.Pressed(pixelgl.KeyRightControl)) {
		eventChannel <- UndoEvent
		return // No next event!
	}

	// Multikey select
	if win.Pressed(pixelgl.KeyLeftShift) || win.Pressed(pixelgl.KeyRightShift) {
		keyWasPressedThisCycle = true
	}

	// Individual keys
	for _, combo := range keyCombos {
		if win.Pressed(combo.Button) {
			keyWasPressedThisCycle = true

			if !keyHasFired[combo.Name] {
				eventChannel <- combo.Event
				keyHasFired[combo.Name] = true
			}
		}
	}

	// Calculate variables
	allKeysReleased := !keyWasPressedThisCycle
	selectionHasBeenMade := len(keyHasFired) > 0

	// Firing next event
	if allKeysReleased && selectionHasBeenMade {
		eventChannel <- NextEvent
		keyHasFired = map[string]bool{}
	}
}

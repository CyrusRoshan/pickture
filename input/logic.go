package input

import (
	"os"

	"github.com/gotk3/gotk3/gdk"
)

// Used to not fire multiple times in one go
var keyHasFired = map[uint]bool{}
var undoHasFired = false

const cmdKeyLeft = 65511
const cmdKeyRight = 65512

func handleKeyChange() {
	var keyWasPressedThisCycle = false

	cmdPressed := (keyIsPressed[gdk.KEY_Super_L] ||
		keyIsPressed[gdk.KEY_Super_R] ||
		keyIsPressed[cmdKeyLeft] ||
		keyIsPressed[cmdKeyRight] ||
		keyIsPressed[gdk.KEY_Control_L] ||
		keyIsPressed[gdk.KEY_Control_R])

	// Quit
	if cmdPressed && keyIsPressed[gdk.KEY_Q] {
		os.Exit(0)
	}

	// Undo event
	if keyIsPressed[gdk.KEY_Z] && cmdPressed {
		if !undoHasFired {
			inputEvents <- UndoEvent
			undoHasFired = true
			return // No next event!
		}
	} else { // Reset undo upon key release
		undoHasFired = false
	}

	// Multikey select
	if keyIsPressed[gdk.KEY_Shift_L] || keyIsPressed[gdk.KEY_Shift_R] {
		keyWasPressedThisCycle = true
	}

	// Individual keys
	for _, combo := range letterKeyEvents {
		if keyIsPressed[combo.Key] {
			keyWasPressedThisCycle = true

			if !keyHasFired[combo.Key] {
				inputEvents <- combo.Event
				keyHasFired[combo.Key] = true
			}
		}
	}

	// Calculate variables
	allKeysReleased := !keyWasPressedThisCycle
	selectionHasBeenMade := len(keyHasFired) > 0

	// Firing next event
	if allKeysReleased && selectionHasBeenMade {
		inputEvents <- NextEvent
		keyHasFired = map[uint]bool{}
	}
}

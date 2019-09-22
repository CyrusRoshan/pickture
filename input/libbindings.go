package input

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

var keyIsPressed = map[uint]bool{}

func BindKeyPressEvents(win *gtk.Window) chan InputEvent {
	var nativeKeyPressEvents = make(chan bool, 3)

	// Helper so we can handle events with knowledge of key value and keypress
	handleGTKKeyEvent := func(event *gdk.Event, pressed bool) {
		key := gdk.EventKeyNewFromEvent(event)
		keyIsPressed[gdk.KeyvalToUpper(key.KeyVal())] = pressed
		nativeKeyPressEvents <- true
	}

	// Handle events serially thanks to the chan
	go func() {
		for {
			<-nativeKeyPressEvents
			handleKeyChange()
		}
	}()

	win.Connect("key-press-event", func(_ *gtk.Window, ev *gdk.Event) {
		handleGTKKeyEvent(ev, true)
	})
	win.Connect("key-release-event", func(_ *gtk.Window, ev *gdk.Event) {
		handleGTKKeyEvent(ev, false)
	})

	return inputEvents
}

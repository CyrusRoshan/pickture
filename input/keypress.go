package input

import (
	"github.com/faiface/pixel/pixelgl"
)

var optionHolder KeyPressOptions

type KeyPressOptions struct {
	Next func()
	Undo func()
	Q    func()
	W    func()
	E    func()
	A    func()
	S    func()
	D    func()
}

func AddKeyPressFunctions(options KeyPressOptions) {
	optionHolder = options
}

// Used to not fire multiple times in one go
var hasFired map[string]bool

func CalculateKeyPressChanges(win *pixelgl.Window) {
	var (
		multiKeySelect = false
		anyKeyPressed  = false
	)

	if win.Pressed(pixelgl.KeyZ) &&
		(win.Pressed(pixelgl.KeyLeftSuper) ||
			win.Pressed(pixelgl.KeyRightSuper) ||
			win.Pressed(pixelgl.KeyLeftControl) ||
			win.Pressed(pixelgl.KeyRightControl)) {
		optionHolder.Undo()
		return
	}

	if win.Pressed(pixelgl.KeyLeftShift) ||
		win.Pressed(pixelgl.KeyRightShift) {
		multiKeySelect = true
	}

	if win.Pressed(pixelgl.KeyQ) && !hasFired["Q"] {
		optionHolder.Q()
		hasFired["Q"] = true
	}
	if win.Pressed(pixelgl.KeyW) && !hasFired["W"] {
		optionHolder.W()
		hasFired["W"] = true
	}
	if win.Pressed(pixelgl.KeyE) && !hasFired["E"] {
		optionHolder.E()
		hasFired["E"] = true
	}
	if win.Pressed(pixelgl.KeyA) && !hasFired["A"] {
		optionHolder.A()
		hasFired["A"] = true
	}
	if win.Pressed(pixelgl.KeyS) && !hasFired["S"] {
		optionHolder.S()
		hasFired["S"] = true
	}
	if win.Pressed(pixelgl.KeyD) && !hasFired["D"] {
		optionHolder.D()
		hasFired["D"] = true
	}

	if !multiKeySelect && anyKeyPressed {
		optionHolder.Next()
		hasFired = map[string]bool{}
	}
}

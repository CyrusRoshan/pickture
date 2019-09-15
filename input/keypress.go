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
var hasFired = map[string]bool{}

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
		if optionHolder.Undo != nil {
			optionHolder.Undo()
		}
		return
	}

	if win.Pressed(pixelgl.KeyLeftShift) ||
		win.Pressed(pixelgl.KeyRightShift) {
		multiKeySelect = true
	}

	if win.Pressed(pixelgl.KeyQ) && !hasFired["Q"] {
		if optionHolder.Q != nil {
			optionHolder.Q()
		}
		anyKeyPressed = true
		hasFired["Q"] = true
	}
	if win.Pressed(pixelgl.KeyW) && !hasFired["W"] {
		if optionHolder.W != nil {
			optionHolder.W()
		}
		anyKeyPressed = true
		hasFired["W"] = true
	}
	if win.Pressed(pixelgl.KeyE) && !hasFired["E"] {
		if optionHolder.E != nil {
			optionHolder.E()
		}
		anyKeyPressed = true
		hasFired["E"] = true
	}
	if win.Pressed(pixelgl.KeyA) && !hasFired["A"] {
		if optionHolder.A != nil {
			optionHolder.A()
		}
		anyKeyPressed = true
		hasFired["A"] = true
	}
	if win.Pressed(pixelgl.KeyS) && !hasFired["S"] {
		if optionHolder.S != nil {
			optionHolder.S()
		}
		anyKeyPressed = true
		hasFired["S"] = true
	}
	if win.Pressed(pixelgl.KeyD) && !hasFired["D"] {
		if optionHolder.D != nil {
			optionHolder.D()
		}
		anyKeyPressed = true
		hasFired["D"] = true
	}

	if !multiKeySelect && anyKeyPressed {
		if optionHolder.Next != nil {
			optionHolder.Next()
		}
		hasFired = map[string]bool{}
	}
}

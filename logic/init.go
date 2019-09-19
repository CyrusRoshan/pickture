package logic

import (
	"github.com/CyrusRoshan/pickture/input"
)

type InitProperties struct {
	InputPath           string
	AOutputPath         string
	DOutputPath         string
	DisableUniqueSuffix bool
}

var props InitProperties

func Init(p InitProperties) {
	props = p

	// Initialize
	State.initialize()
	inputEvents := input.GetKeyPressEvents()
	go handleInputEvents(inputEvents)
}

package logic

import (
	"github.com/CyrusRoshan/pickture/input"
)

type InitProperties struct {
	// Input and output folder paths
	InputPath   string
	AOutputPath string
	SOutputPath string
	DOutputPath string
	QOutputPath string
	WOutputPath string
	EOutputPath string

	// The suffix goes on the end of output files to avoid collisions when
	// multiple input files in different nested folders have the same filename
	DisableUniqueSuffix bool

	InputEvents chan input.InputEvent

	// Useful for only re-rendering on state change
	OnChange func()
}

var props InitProperties

func Init(p InitProperties) {
	props = p
	State.initialize()
	p.OnChange()
	go handleInputEvents(p.InputEvents, p.OnChange)
}

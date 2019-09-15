package ui

import (
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/faiface/pixel/pixelgl"
)

const (
	Title  = "pickture"
	MaxFPS = 30
)

func Render() {
	window, err := BuildWindow(Title)
	utils.PanicIfErr(err)

	// Initial setup
	Setup(window)

	// While running...
	for !window.Closed() {
		LimitFPS(30, func() {
			ShowFPSInTitle(Title, window)
			Changes(window)
			window.Update()
		})
	}

	// Clean up before exiting
	CleanUp(window)
}

func Setup(window *pixelgl.Window) {
	getNewInfo()
	SetBindings()
}

func Changes(window *pixelgl.Window) {
	RenderCurrentState(window)
}

func CleanUp(window *pixelgl.Window) {

}

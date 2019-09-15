package ui

import (
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	Title  = "pickture"
	MaxFPS = 30
)

var (
	window *pixelgl.Window
)

func Render() {
	window, err := BuildWindow(Title)
	utils.PanicIfErr(err)

	// Initial setup
	Setup()

	// While running...
	for !window.Closed() {
		LimitFPS(30, func() {
			ShowFPSInTitle(Title, window)
			Changes()
			window.Update()
		})
	}

	// Clean up before exiting
	CleanUp()
}

func Setup() {
	window.Clear(colornames.Skyblue)
	QButton.Draw(window, pixel.IM.Scaled(pixel.V(0, 0), 2).Moved(window.Bounds().Center()))
}

func Changes() {

}

func CleanUp() {

}

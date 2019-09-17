package main

import (
	"github.com/CyrusRoshan/pickture/input"
	"github.com/CyrusRoshan/pickture/logic"
	"github.com/CyrusRoshan/pickture/ui"
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/alexflint/go-arg"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(Render)
}

const (
	Title  = "pickture"
	MaxFPS = 30
)

func Render() {
	window, err := ui.BuildWindow(Title)
	utils.PanicIfErr(err)

	// Initial setup
	Setup(window)

	// While running...
	for !window.Closed() {
		ui.LimitFPS(30, func() {
			ui.ShowFPSInTitle(Title, window)
			RenderChanges(window)
			window.Update()
		})
	}

	// Clean up before exiting
	CleanUp(window)
}

func Setup(win *pixelgl.Window) {
	var args = struct {
		Input  string `arg:"positional" help:"test"`
		Output string `arg:"positional" help:"test"`
	}{
		Input:  "./input",
		Output: "./output",
	}
	arg.MustParse(&args)

	logic.Init(logic.InitProperties{
		InputPath:   args.Input,
		AOutputPath: args.Output + "/a",
		DOutputPath: args.Output + "/d",
	})
}

func RenderChanges(win *pixelgl.Window) {
	// Get changes to previous state
	input.CalculateKeyPressChanges(win)

	// Draw current state
	win.Clear(colornames.Black) // Start with black background
	if currFile := logic.CurrentFile(); currFile != nil {
		ui.DrawBackgroundImage(win, *currFile) // Draw current image
	}
	ui.DrawButtons(win)
	ui.DrawImageCount(win, logic.GetImageCount())
}

func CleanUp(win *pixelgl.Window) {

}

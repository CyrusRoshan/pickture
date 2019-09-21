package main

import (
	"image"
	"path/filepath"

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
		Input               string `arg:"positional" help:"Input folder (where all input files are stored in)"`
		Output              string `arg:"positional" help:"Output folder (where the sorted folders and their contents will go)"`
		DisableUniqueSuffix bool   `arg:"--nosuffix" help:"disable unique suffix for images. Warning: may result in errors if nested input folders have files with conflicting names."`
	}{
		Input:               "./input",
		Output:              "./output",
		DisableUniqueSuffix: false,
	}
	arg.MustParse(&args)

	getAbsPath := func(path string) string {
		abs, err := filepath.Abs(path)
		utils.PanicIfErr(err, "Could not get absolute path for "+path)
		return abs
	}

	logic.Init(logic.InitProperties{
		InputPath:           getAbsPath(args.Input),
		AOutputPath:         getAbsPath(args.Output + "/A"),
		SOutputPath:         getAbsPath(args.Output + "/S"),
		DOutputPath:         getAbsPath(args.Output + "/D"),
		QOutputPath:         getAbsPath(args.Output + "/Q"),
		WOutputPath:         getAbsPath(args.Output + "/W"),
		EOutputPath:         getAbsPath(args.Output + "/E"),
		DisableUniqueSuffix: args.DisableUniqueSuffix,
	})
}

func RenderChanges(win *pixelgl.Window) {
	// Get changes to previous state
	input.CalculateKeyPressChanges(win)

	// Draw current state
	win.Clear(colornames.Black) // Start with black background

	imageName := "[none]"
	if currFile := logic.State.GetCurrentFile(); currFile != nil {
		var currImg *image.Image

		utils.LogTimeSpent(func() {
			currImg = logic.State.GetCurrentImage()
		}, "getting image")
		utils.LogTimeSpent(func() {
			ui.DrawBackgroundImage(win, currImg)
		}, "drawing background image")
		imageName = currFile.Info.Name()
	}
	ui.DrawButtons(win)
	ui.DrawImageInfo(win, logic.State.GetImageCount(), imageName)
}

func CleanUp(win *pixelgl.Window) {

}

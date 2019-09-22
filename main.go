package main

import (
	"fmt"
	"path/filepath"

	"github.com/CyrusRoshan/pickture/input"
	"github.com/CyrusRoshan/pickture/logic"
	"github.com/CyrusRoshan/pickture/ui"
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/alexflint/go-arg"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	Render()
}

const (
	Title  = "pickture"
	MaxFPS = 30
)

func Render() {
	// Init GTK
	gtk.Init(nil)

	// Build window
	window, err := ui.BuildWindow(Title)
	utils.PanicIfErr(err, "Could not build window")

	// Initial setup
	fmt.Println("INIT")
	SetupInternals(window,
		func() {
			RenderChanges(window)
		},
	)

	fmt.Println("NEWBOX")
	container, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	utils.PanicIfErr(err, "error creating box")
	buttonGrid := ui.GetButtonGrid()
	container.Add(buttonGrid)
	window.Add(container)

	// Execute and block main thread
	gtk.Main()
}

func SetupInternals(window *gtk.Window, OnChange func()) {
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
		InputPath:   getAbsPath(args.Input),
		AOutputPath: getAbsPath(args.Output + "/A"),
		SOutputPath: getAbsPath(args.Output + "/S"),
		DOutputPath: getAbsPath(args.Output + "/D"),
		QOutputPath: getAbsPath(args.Output + "/Q"),
		WOutputPath: getAbsPath(args.Output + "/W"),
		EOutputPath: getAbsPath(args.Output + "/E"),

		DisableUniqueSuffix: args.DisableUniqueSuffix,

		InputEvents: input.BindKeyPressEvents(window),

		OnChange: OnChange,
	})
}

func RenderChanges(win *gtk.Window) {
	fmt.Println("RENDERING CHANGES!")
	return
	// container.PackStart(child, expand, fill, padding)

	// container.PackEnd(child, expand, fill, padding)

	// imageName := "[none]"
	// if currFile := logic.State.GetCurrentFile(); currFile != nil {
	// 	var currImg *image.Image

	// 	utils.LogTimeSpent(func() {
	// 		currImg = logic.State.GetCurrentImage()
	// 	}, "getting image")
	// 	utils.LogTimeSpent(func() {
	// 		ui.DrawBackgroundImage(win, currImg)
	// 	}, "drawing background image")
	// 	imageName = currFile.Info.Name()
	// }
	// ui.DrawButtons(win)
	// ui.DrawImageInfo(win, logic.State.GetImageCount(), imageName)
}

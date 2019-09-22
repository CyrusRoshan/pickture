package main

import (
	"path/filepath"

	"github.com/CyrusRoshan/pickture/input"
	"github.com/CyrusRoshan/pickture/logic"
	"github.com/CyrusRoshan/pickture/ui"
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/alexflint/go-arg"
	"github.com/gotk3/gotk3/gtk"
)

const (
	Title  = "pickture"
	MaxFPS = 30
)

func main() {
	// Initialize GUI library
	gtk.Init(nil)
	window, err := ui.BuildWindow(Title)
	utils.PanicIfErr(err, "Could not build window")

	// Initialize pickture logic
	rootWidget, updateRender := ui.Root()
	SetupInternals(window, updateRender)

	// Hook in and render the UI
	window.Add(rootWidget)
	window.ShowAll()

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

package ui

import (
	"log"

	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gobuffalo/packr"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

var (
	QButton *gdk.Pixbuf
	WButton *gdk.Pixbuf
	EButton *gdk.Pixbuf
	AButton *gdk.Pixbuf
	SButton *gdk.Pixbuf
	DButton *gdk.Pixbuf
)

func init() {
	packBox := packr.NewBox("../assets/raw") // Helps bundle files to binary
	var getButtonFromAssets = func(filename string) *gdk.Pixbuf {
		fullFilename := filename + "_button.png"

		pixbuf, err := PixbufFromFile(fullFilename, &packBox)
		utils.PanicIfErr(err)
		return pixbuf
	}

	QButton = getButtonFromAssets("q")
	WButton = getButtonFromAssets("w")
	EButton = getButtonFromAssets("e")
	AButton = getButtonFromAssets("a")
	SButton = getButtonFromAssets("s")
	DButton = getButtonFromAssets("d")
}

func ButtonsHolder() *gtk.Grid {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	insertBtn, err := gtk.ButtonNewWithLabel("Add a label")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	removeBtn, err := gtk.ButtonNewWithLabel("Remove a label")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	grid.Add(insertBtn)
	grid.Attach(removeBtn, 1, 1, 1, 1)

	return grid

	// width := QButton.Frame().W()
	// height := QButton.Frame().H()
	// spacer := float64(5)

	// xSpace := (width + spacer)
	// ySpace := (height + spacer)

	// unscaledYShift := -(win.Bounds().H() / 2)
	// scaledYShift := ySpace + height/2 + spacer

	// var drawButton = func(button *pixel.Sprite, x float64, y float64) {
	// 	button.Draw(win,
	// 		pixel.IM.
	// 			Moved(pixel.V(x, y)).
	// 			Moved(pixel.V(0, scaledYShift)).
	// 			Scaled(pixel.V(0, 0), 2).
	// 			Moved(pixel.V(0, unscaledYShift)).
	// 			Moved(win.Bounds().Center()))
	// }

	// drawButton(QButton, -xSpace, 0)
	// drawButton(WButton, 0, 0)
	// drawButton(EButton, xSpace, 0)

	// drawButton(AButton, -xSpace, -ySpace)
	// drawButton(SButton, 0, -ySpace)
	// drawButton(DButton, xSpace, -ySpace)
}

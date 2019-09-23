package widgets

import (
	"fmt"
	"path/filepath"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gobuffalo/packr"

	"github.com/gotk3/gotk3/gdk"
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
	assetsPath, err := filepath.Abs("./assets/raw")
	utils.PanicIfErr(err)
	packBox := packr.NewBox(assetsPath) // Helps bundle files to binary

	var getButtonFromAssets = func(filename string) *gdk.Pixbuf {
		fullFilename := filename + "_button.png"

		pixbuf, err := files.PixbufFromFile(fullFilename, &packBox)
		utils.PanicIfErr(
			err,
			fmt.Sprintf("Loading pixbuf from file: %s", fullFilename),
		)
		return pixbuf
	}

	QButton = getButtonFromAssets("q")
	WButton = getButtonFromAssets("w")
	EButton = getButtonFromAssets("e")
	AButton = getButtonFromAssets("a")
	SButton = getButtonFromAssets("s")
	DButton = getButtonFromAssets("d")
}

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

package ui

import (
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/gobuffalo/packr"
)

var (
	QButton *pixel.Sprite
	WButton *pixel.Sprite
	EButton *pixel.Sprite
	AButton *pixel.Sprite
	SButton *pixel.Sprite
	DButton *pixel.Sprite
)

func init() {
	packBox := packr.NewBox("../assets/raw") // Helps bundle files to binary
	var getButtonFromAssets = func(filename string) *pixel.Sprite {
		fullFilename := filename + "_button.png"

		sprite, err := SpriteFromFile(fullFilename, &packBox)
		utils.PanicIfErr(err)
		return sprite
	}

	QButton = getButtonFromAssets("q")
	WButton = getButtonFromAssets("w")
	EButton = getButtonFromAssets("e")
	AButton = getButtonFromAssets("a")
	SButton = getButtonFromAssets("s")
	DButton = getButtonFromAssets("d")
}

func DrawButtons(win *pixelgl.Window) {
	width := QButton.Frame().W()
	height := QButton.Frame().H()
	spacer := float64(5)

	xSpace := (width + spacer)
	ySpace := (height + spacer)

	unscaledYShift := -(win.Bounds().H() / 2)
	scaledYShift := ySpace + height/2 + spacer

	var drawButton = func(button *pixel.Sprite, x float64, y float64) {
		button.Draw(win,
			pixel.IM.
				Moved(pixel.V(x, y)).
				Moved(pixel.V(0, scaledYShift)).
				Scaled(pixel.V(0, 0), 2).
				Moved(pixel.V(0, unscaledYShift)).
				Moved(win.Bounds().Center()))
	}

	drawButton(QButton, -xSpace, 0)
	drawButton(WButton, 0, 0)
	drawButton(EButton, xSpace, 0)

	drawButton(AButton, -xSpace, -ySpace)
	drawButton(SButton, 0, -ySpace)
	drawButton(DButton, xSpace, -ySpace)
}

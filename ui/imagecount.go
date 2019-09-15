package ui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

func DrawImageCount(win *pixelgl.Window, imageCount int) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(0, 0), basicAtlas)
	fmt.Fprintf(basicTxt, "Images: (%d)", imageCount)

	basicTxt.Draw(win, pixel.IM.
		Moved(win.Bounds().Center()).
		Moved(pixel.V(
			-win.Bounds().W()/2+10,
			win.Bounds().H()/2-(basicTxt.Bounds().H()+10))))
}

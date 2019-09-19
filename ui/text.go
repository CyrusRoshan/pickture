package ui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

func DrawImageInfo(win *pixelgl.Window, count int, name string) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// Write text
	imageCount := text.New(pixel.V(0, 0), basicAtlas)
	fmt.Fprintf(imageCount, "Images: (%d)", count)
	imageName := text.New(pixel.V(0, 0), basicAtlas)
	fmt.Fprintf(imageName, "Name: (%s)", name)

	// Do calculations
	const horizontalPadding = float64(10)
	horizontalPosition := horizontalPadding - win.Bounds().W()/2

	const verticalPadding = float64(10)
	verticalPosition := win.Bounds().H()/2 - verticalPadding

	textHeight := imageCount.Bounds().H()
	const elementSpacing = float64(5)

	// Draw text
	imageCount.Draw(win, pixel.IM.
		Moved(win.Bounds().Center()).
		Moved(pixel.V(
			horizontalPosition,
			verticalPosition-textHeight,
		)),
	)

	imageName.Draw(win, pixel.IM.
		Moved(win.Bounds().Center()).
		Moved(pixel.V(
			horizontalPosition,
			verticalPosition-(textHeight*2+elementSpacing),
		)),
	)
}

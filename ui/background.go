package ui

import (
	"math"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func imageScalingRatio(winBounds pixel.Rect, imageBounds pixel.Rect) float64 {
	wRatio := winBounds.W() / imageBounds.W()
	hRatio := winBounds.H() / imageBounds.H()

	smallestRatio := math.Min(wRatio, hRatio)
	return smallestRatio
}

var imageCache = map[string]*pixel.Sprite{}

func DrawBackgroundImage(win *pixelgl.Window, image files.File) {
	imageSprite, ok := imageCache[image.Path]
	if !ok {
		var err error
		imageSprite, err = SpriteFromFile(image.Path)
		utils.PanicIfErr(err)

		imageCache[image.Path] = imageSprite
	}

	scaleRatio := imageScalingRatio(win.Bounds(), imageSprite.Frame())
	imageSprite.Draw(win, pixel.IM.
		Scaled(pixel.ZV, scaleRatio).
		Moved(win.Bounds().Center()))
}

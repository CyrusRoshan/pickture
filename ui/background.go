package ui

import (
	"image"
	"math"

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

func DrawBackgroundImage(win *pixelgl.Window, img *image.Image) {
	var ps *pixel.Sprite
	utils.LogTimeSpent(func() {
		pd := pixel.PictureDataFromImage(*img)
		ps = pixel.NewSprite(pd, pd.Bounds())
	}, "converting to sprite")

	utils.LogTimeSpent(func() {
		scaleRatio := imageScalingRatio(
			win.Bounds(),
			ps.Frame(),
		)
		ps.Draw(win,
			pixel.IM.
				Scaled(pixel.ZV, scaleRatio).
				Moved(win.Bounds().Center()),
		)
	}, "actually drawing and scaling")
}

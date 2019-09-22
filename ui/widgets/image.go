package widgets

import (
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type TitledImageHolderState struct {
	Title       string
	ImagePixbuf *gdk.Pixbuf
}

func TitledImageHolderWidget() *UpdaterWidget {
	frame, err := gtk.FrameNew(DefaultText)
	utils.PanicIfErr(err)

	img, err := gtk.ImageNew()
	utils.PanicIfErr(err)

	gtkWidget := &frame.Container.Widget

	// Add update func, which can access the pointers we just created
	var updateFunc = func(state interface{}) {
		s, ok := state.(TitledImageHolderState)
		if !ok {
			panic("Error converting state from interface")
		}

		_, err := glib.IdleAdd(img.SetFromPixbuf, s.ImagePixbuf)
		utils.PanicIfErr(err)

		_, err = glib.IdleAdd(frame.SetLabel, s.Title)
		utils.PanicIfErr(err)
	}
	updaterWidget := NewUpdaterWidget(gtkWidget, updateFunc)
	return &updaterWidget
}

// func imageScalingRatio(winBounds pixel.Rect, imageBounds pixel.Rect) float64 {
// 	wRatio := winBounds.W() / imageBounds.W()
// 	hRatio := winBounds.H() / imageBounds.H()

// 	smallestRatio := math.Min(wRatio, hRatio)
// 	return smallestRatio
// }
// func DrawBackgroundImage(win *pixelgl.Window, img *image.Image) {
// 	var ps *pixel.Sprite
// 	utils.LogTimeSpent(func() {
// 		pd := pixel.PictureDataFromImage(*img)
// 		ps = pixel.NewSprite(pd, pd.Bounds())
// 	}, "converting to sprite")

// 	utils.LogTimeSpent(func() {
// 		scaleRatio := imageScalingRatio(
// 			win.Bounds(),
// 			ps.Frame(),
// 		)
// 		ps.Draw(win,
// 			pixel.IM.
// 				Scaled(pixel.ZV, scaleRatio).
// 				Moved(win.Bounds().Center()),
// 		)
// 	}, "actually drawing and scaling")
// }

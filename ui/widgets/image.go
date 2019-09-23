package widgets

import (
	"math"

	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type TitledImageHolderState struct {
	Title       string
	ImagePixbuf *gdk.Pixbuf
}

func TitledImageHolderWidget(parent *gtk.Widget) *UpdaterWidget {
	frame, err := gtk.FrameNew(DefaultText)
	utils.PanicIfErr(err)
	frame.SetHExpand(true)
	frame.SetVExpand(true)

	img, err := gtk.ImageNew()
	utils.PanicIfErr(err)
	img.SetHExpand(true)
	img.SetVExpand(true)
	frame.Add(img)

	gtkWidget := &frame.Container.Widget

	// Add update func, which can access the pointers we just created
	var updateFunc = func(state interface{}) {
		s, ok := state.(TitledImageHolderState)
		if !ok {
			panic("Error converting state from interface")
		}

		var pixbuf *gdk.Pixbuf
		if s.ImagePixbuf != nil {
			w, h := scaleImage(
				parent.GetAllocatedWidth(),
				parent.GetAllocatedHeight(),
				s.ImagePixbuf.GetWidth(),
				s.ImagePixbuf.GetHeight(),
				0.9,
			)

			pixbuf, err = s.ImagePixbuf.ScaleSimple(w, h, gdk.INTERP_TILES)
			utils.PanicIfErr(err, "error scaling pixbuf")
		}

		_, err = glib.IdleAdd(img.SetFromPixbuf, pixbuf)
		utils.PanicIfErr(err)

		_, err = glib.IdleAdd(frame.SetLabel, s.Title)
		utils.PanicIfErr(err)
	}
	updaterWidget := NewUpdaterWidget(gtkWidget, updateFunc)
	return &updaterWidget
}

func scaleImage(fitWidth, fitHeight, imageWidth, imageHeight int, ratio float64) (width, height int) {
	ratio *= imageScalingRatio(fitWidth, fitHeight, imageWidth, imageHeight)
	return int(float64(imageWidth) * ratio), int(float64(imageHeight) * ratio)
}

func imageScalingRatio(fitWidth, fitHeight, imageWidth, imageHeight int) float64 {
	wRatio := float64(fitWidth) / float64(imageWidth)
	hRatio := float64(fitHeight) / float64(imageHeight)

	smallestRatio := math.Min(wRatio, hRatio)
	return smallestRatio
}

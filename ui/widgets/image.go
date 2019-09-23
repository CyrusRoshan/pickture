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

	var renderNewPixbuf = func(pixbuf *gdk.Pixbuf) error {
		if pixbuf != nil {
			w, h := scaleImage(
				parent.GetAllocatedWidth(),
				parent.GetAllocatedHeight(),
				pixbuf.GetWidth(),
				pixbuf.GetHeight(),
				0.9,
			)

			pixbuf, err = pixbuf.ScaleSimple(w, h, gdk.INTERP_TILES)
			if err != nil {
				return err
			}
		}

		_, err = glib.IdleAdd(img.SetFromPixbuf, pixbuf)
		if err != nil {
			return err
		}
		return nil
	}

	oldPWidth := parent.GetAllocatedWidth()
	oldPHeight := parent.GetAllocatedHeight()
	var resizeImage = func(pixbuf *gdk.Pixbuf) {
		// Check if we need to resize
		pWidth := parent.GetAllocatedWidth()
		pHeight := parent.GetAllocatedHeight()
		if pWidth == oldPWidth && pHeight == oldPHeight {
			return
		}
		oldPWidth, oldPHeight = pWidth, pHeight

		err := renderNewPixbuf(pixbuf)
		utils.PanicIfErr(err, "error rendering pixbuf")
	}

	var lastPixbuf *gdk.Pixbuf
	img.Connect("draw", func() { resizeImage(lastPixbuf) })

	// Add update func, which can access the pointers we just created
	var updateFunc = func(state interface{}) {
		s, ok := state.(TitledImageHolderState)
		if !ok {
			panic("Error converting state from interface")
		}

		// Sorry this is basically state abuse
		lastPixbuf = s.ImagePixbuf

		err := renderNewPixbuf(s.ImagePixbuf)
		utils.PanicIfErr(err, "error rendering pixbuf")

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

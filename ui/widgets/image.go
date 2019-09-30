package widgets

import (
	"math"

	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gotk3/gotk3/cairo"
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

	drawing, err := gtk.DrawingAreaNew()
	utils.PanicIfErr(err)
	frame.Add(drawing)

	gtkWidget := &frame.Container.Widget

	var resizeImage = func(pixbuf *gdk.Pixbuf, maxWidth, maxHeight int) (*gdk.Pixbuf, int, int) {
		width, height := scaleImage(
			maxWidth,
			maxHeight,
			pixbuf.GetWidth(),
			pixbuf.GetHeight(),
			1,
		)

		pixbuf, err = pixbuf.ScaleSimple(width, height, gdk.INTERP_TILES)
		utils.PanicIfErr(err, "error resizing pixbuf")
		return pixbuf, maxWidth - width, maxHeight - height
	}

	var lastPixbuf *gdk.Pixbuf
	drawing.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		height := da.GetAllocatedHeight()
		width := da.GetAllocatedWidth()
		pixbuf, wdiff, hdiff := resizeImage(lastPixbuf, width, height)

		cr.Rectangle(0, 0, float64(width), float64(height))
		gtk.GdkCairoSetSourcePixBuf(cr, pixbuf,
			float64(wdiff/2), float64(hdiff/2))
		cr.Fill()
	})

	// Add update func, which can access the pointers we just created
	var updateFunc = func(state interface{}) {
		s, ok := state.(TitledImageHolderState)
		if !ok {
			panic("Error converting state from interface")
		}

		lastPixbuf = s.ImagePixbuf
		drawing.QueueDraw()

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

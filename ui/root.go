package ui

import (
	"github.com/CyrusRoshan/pickture/ui/widgets"

	"github.com/gotk3/gotk3/gtk"

	"github.com/CyrusRoshan/pickture/logic"
	"github.com/CyrusRoshan/pickture/utils"
)

func Root() (rootWidget *gtk.Widget, onUpdate func()) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	utils.PanicIfErr(err, "error creating box")

	// box.Add(ButtonsHolder())
	// imageHolder := widgets.TitledImageHolderWidget()

	imageInfo := widgets.ImageInfoHolderWidget()
	box.Add(imageInfo.InnerWidget)

	onUpdate = func() {
		imageName := "[none]"
		if currFile := logic.State.GetCurrentFile(); currFile != nil {
			imageName = currFile.Info.Name()

			// var currImg *image.Image
			// currImg = logic.State.GetCurrentImage()

			// DrawBackgroundImage(win, currImg)
		}

		imageInfo.Update(widgets.ImageInfoHolderState{
			Count: logic.State.GetImageCount(),
			Name:  imageName,
		})
	}

	return &box.Container.Widget, onUpdate
}

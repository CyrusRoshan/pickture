package ui

import (
	"github.com/CyrusRoshan/pickture/ui/widgets"

	"github.com/gotk3/gotk3/gtk"

	"github.com/CyrusRoshan/pickture/logic"
	"github.com/CyrusRoshan/pickture/utils"
)

func Root(window *gtk.Window) (rootWidget *gtk.Widget, onUpdate func()) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	utils.PanicIfErr(err, "error creating box")

	// box.Add(ButtonsHolder())

	imageHolder := widgets.TitledImageHolderWidget(&window.Container.Widget)
	box.Add(imageHolder.InnerWidget)

	imageInfo := widgets.ImageInfoHolderWidget()
	box.Add(imageInfo.InnerWidget)

	onUpdate = func() {
		imageName := "[none]"
		if currFile := logic.State.GetCurrentFile(); currFile != nil {
			imageName = currFile.Info.Name()
		}

		imageHolder.Update(widgets.TitledImageHolderState{
			Title:       imageName,
			ImagePixbuf: logic.State.GetCurrentImage(),
		})

		imageInfo.Update(widgets.ImageInfoHolderState{
			Index: logic.State.GetCurrentImageIndex(),
			Count: logic.State.GetImageCount(),
			Name:  imageName,
		})
	}

	return &box.Container.Widget, onUpdate
}

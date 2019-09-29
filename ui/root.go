package ui

import (
	"github.com/CyrusRoshan/pickture/ui/widgets"

	"github.com/gotk3/gotk3/gtk"

	"github.com/CyrusRoshan/pickture/logic"
	"github.com/CyrusRoshan/pickture/utils"
)

func Root(window *gtk.Window) (rootWidget *gtk.Widget, onUpdate func()) {
	addCSS()

	// Create widgets
	imageInfo := widgets.ImageInfoHolderWidget()
	imageHolder := widgets.TitledImageHolderWidget(&window.Container.Widget)

	// Create container elements
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	utils.PanicIfErr(err, "error creating box")
	{

		leftSide, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
		utils.PanicIfErr(err, "error creating box")
		box.Add(leftSide)

		topLeft, err := gtk.FrameNew("Image Info:")
		utils.PanicIfErr(err)
		leftSide.Add(topLeft)
		bottomLeft, err := gtk.FrameNew("Sort Folder Info:")
		utils.PanicIfErr(err)
		leftSide.Add(bottomLeft)

		{ // Add widgets to their containers
			topLeft.Add(imageInfo.InnerWidget)
			box.Add(imageHolder.InnerWidget)
		}
	}

	// info to have: ---
	// total image count (curr/number)
	// progress bar showing %done
	// current image full path
	// current image metadata (day/time)
	// current image metadata camera, lens info
	//
	// binding info: ---
	// q w e a s d nicknames
	// q w e a s d image counts
	// q w e a s d (show show last selection highlighted, show selected keys
	// with border, useful for multi-select)
	//
	//
	// Example: ----
	// [ info   ] [ picture      ]
	// [        ] [              ]
	// [        ] [              ]
	// [ bindings][              ]
	// [ q w e  ] [              ]
	// [ a s d  ] [              ]

	// Set state changes for widgets
	onUpdate = func() {
		imageName := "[none]"
		imagePath := "[none]"
		if currFile := logic.State.GetCurrentFile(); currFile != nil {
			imageName = currFile.Info.Name()
			imagePath = currFile.Path
		}

		imageHolder.Update(widgets.TitledImageHolderState{
			Title:       imageName,
			ImagePixbuf: logic.State.GetCurrentImage(),
		})

		imageInfo.Update(widgets.ImageInfoHolderState{
			Index: logic.State.GetCurrentImageIndex(),
			Count: logic.State.GetImageCount(),
			Name:  imageName,
			Path:  imagePath,
		})
	}

	return &box.Container.Widget, onUpdate
}

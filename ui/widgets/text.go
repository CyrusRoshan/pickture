package widgets

import (
	"fmt"

	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type ImageInfoHolderState struct {
	Index int
	Count int
	Name  string
	Path  string
}

func ImageInfoHolderWidget() *UpdaterWidget {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	utils.PanicIfErr(err)

	newLabel := func() *gtk.Label {
		label, err := gtk.LabelNew(DefaultText)
		utils.PanicIfErr(err)
		label.SetWidthChars(30)
		label.SetMaxWidthChars(30)
		label.SetLineWrap(true)

		box.Add(label)
		return label
	}

	imageCountLabel := newLabel()
	imagePathLabel := newLabel()
	// current image metadata (day/time)
	// current image metadata camera, lens info
	progressBar, err := gtk.ProgressBarNew()
	utils.PanicIfErr(err)
	box.Add(progressBar)

	gtkWidget := &box.Container.Widget

	// Add update func, which can access the pointers we just created
	var updateFunc = func(state interface{}) {
		s, ok := state.(ImageInfoHolderState)
		if !ok {
			panic("Error converting state from interface")
		}

		var countText string
		if s.Index >= s.Count {
			countText = fmt.Sprintf(
				"Images: (%d/%d)",
				s.Count,
				s.Count,
			)
		} else {
			countText = fmt.Sprintf(
				"Images: (%d/%d)",
				s.Index+1,
				s.Count,
			)
		}
		_, err := glib.IdleAdd(imageCountLabel.SetText, countText)
		utils.PanicIfErr(err)

		_, err = glib.IdleAdd(imagePathLabel.SetText, fmt.Sprintf("Path: %s", s.Path))
		utils.PanicIfErr(err)

		progressBar.SetFraction(float64(s.Index) / float64(s.Count))
	}
	updaterWidget := NewUpdaterWidget(gtkWidget, updateFunc)
	return &updaterWidget
}

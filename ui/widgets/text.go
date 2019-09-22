package widgets

import (
	"fmt"

	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type ImageInfoHolderState struct {
	Count int
	Name  string
}

func ImageInfoHolderWidget() *UpdaterWidget {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	utils.PanicIfErr(err)

	countLabel, err := gtk.LabelNew(DefaultText)
	utils.PanicIfErr(err)
	box.Add(countLabel)

	nameLabel, err := gtk.LabelNew(DefaultText)
	utils.PanicIfErr(err)
	box.Add(nameLabel)

	gtkWidget := &box.Container.Widget

	// Add update func, which can access the pointers we just created
	var updateFunc = func(state interface{}) {
		s, ok := state.(ImageInfoHolderState)
		if !ok {
			panic("Error converting state from interface")
		}

		_, err := glib.IdleAdd(countLabel.SetText, fmt.Sprintf("Images: (%d)", s.Count))
		utils.PanicIfErr(err)

		_, err = glib.IdleAdd(nameLabel.SetText, fmt.Sprintf("Name: (%s)", s.Name))
		utils.PanicIfErr(err)
	}
	updaterWidget := NewUpdaterWidget(gtkWidget, updateFunc)
	return &updaterWidget
}

package widgets

import (
	"fmt"
	"strings"

	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rwcarlsen/goexif/exif"
)

type ImageInfoHolderState struct {
	Index    int
	Count    int
	Name     string
	Path     string
	ExifData *exif.Exif
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

	newTreeColumn := func(title string, index int) *gtk.TreeViewColumn {
		cellRenderer, err := gtk.CellRendererTextNew()
		utils.PanicIfErr(err)

		column, err := gtk.TreeViewColumnNewWithAttribute(title, cellRenderer, "text", index)
		utils.PanicIfErr(err)
		return column
	}

	// Text labels
	imageCountLabel := newLabel()
	imagePathLabel := newLabel()

	// Metadata box
	metadataView, err := gtk.TreeViewNew()
	metadataView.SetCanFocus(false)
	utils.PanicIfErr(err)
	metadataView.AppendColumn(newTreeColumn("Metadata", 0))
	metadataView.AppendColumn(newTreeColumn("Value", 1))
	box.Add(metadataView)

	// Progress bar
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

		var metadataList *gtk.ListStore
		if s.ExifData != nil {
			metadataList = getMetadataList(s.ExifData)
		}
		_, err = glib.IdleAdd(metadataView.SetModel, metadataList)
		utils.PanicIfErr(err)

		progressBar.SetFraction(float64(s.Index) / float64(s.Count))
	}
	updaterWidget := NewUpdaterWidget(gtkWidget, updateFunc)
	return &updaterWidget
}

func getMetadataList(exifData *exif.Exif) *gtk.ListStore {
	mdls, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)
	utils.PanicIfErr(err)

	// var iterator *gtk.TreeIter
	appendRow := func(key, value string) {
		newRow := mdls.Append()

		err = mdls.SetValue(newRow, 0, key)
		utils.PanicIfErr(err)
		err = mdls.SetValue(newRow, 1, value)
		utils.PanicIfErr(err)
	}

	getField := func(field exif.FieldName) string {
		value, err := exifData.Get(field)
		if err == nil {
			s := value.String()
			return strings.ReplaceAll(s, "\"", "")
		}
		return ""
	}

	camModel := getField(exif.Model)
	if camModel != "" {
		appendRow("Camera Model", camModel)
	}

	appendRow("ISO", getField(exif.ISOSpeedRatings))

	fnumber, err := exifData.Get(exif.FNumber)
	if err == nil {
		numer, denom, err := fnumber.Rat2(0)
		if err == nil {
			simplifiedFloatStr := strings.TrimRight(
				strings.TrimRight(
					fmt.Sprintf("%.2f", float64(numer)/float64(denom)),
					"0"),
				".",
			)
			appendRow("F Number", fmt.Sprintf("f/%s", simplifiedFloatStr))
		}
	}

	focal, err := exifData.Get(exif.FocalLength)
	if err == nil {
		numer, denom, err := focal.Rat2(0)
		if err == nil {
			appendRow("Focal Length", fmt.Sprintf("%.0fmm", float64(numer)/float64(denom)))
		}
	}

	exposure, err := exifData.Get(exif.ExposureTime)
	if err == nil {
		numer, denom, err := exposure.Rat2(0)
		if err == nil {
			numer, denom := utils.SimplifyFraction(int(numer), int(denom))
			var value string
			if denom == 1 {
				value = fmt.Sprintf("%ds", numer)
			} else {
				value = fmt.Sprintf("%d/%ds", numer, denom)
			}
			appendRow("Exposure Time", value)
		}
	}

	tm, err := exifData.DateTime()
	if err == nil {
		appendRow("Date Taken", tm.Format("Monday January 2 2006\n3:04 PM MST"))
	}

	lat, long, err := exifData.LatLong()
	if err == nil {
		appendRow("Coordinates", fmt.Sprintf("Lat: %.5f Long: %.5f", lat, long))
	}
	return mdls
}

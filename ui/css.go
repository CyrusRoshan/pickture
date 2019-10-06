package ui

import (
	"io/ioutil"

	"github.com/CyrusRoshan/pickture/assets"
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func addCSS() {
	// Read file
	f, err := assets.RawAssetsBox.Open("example.css")
	utils.PanicIfErr(err)
	defer f.Close()
	fBytes, err := ioutil.ReadAll(f)
	utils.PanicIfErr(err)

	// Load CSS
	css, err := gtk.CssProviderNew()
	utils.PanicIfErr(err)
	err = css.LoadFromData(string(fBytes))
	utils.PanicIfErr(err)

	// Apply CSS
	defaultScreen, err := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(defaultScreen, css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
}

package ui

import (
	"io/ioutil"
	"path/filepath"

	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gobuffalo/packr"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func addCSS() {
	// Create box
	assetsPath, err := filepath.Abs("./assets/raw")
	utils.PanicIfErr(err)
	packBox := packr.NewBox(assetsPath)

	// Read file
	f, err := packBox.Open("example.css")
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

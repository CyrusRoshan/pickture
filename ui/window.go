package ui

import (
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/gotk3/gotk3/gtk"
)

func BuildWindow(title string) (*gtk.Window, error) {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Get default size
	screen, err := win.GetScreen()
	utils.PanicIfErr(err)
	display, err := screen.GetDisplay()
	utils.PanicIfErr(err)
	monitor, err := display.GetPrimaryMonitor()
	size := monitor.GetGeometry()

	// Set the default window size.
	win.SetDefaultSize(size.GetWidth()*3/4, size.GetHeight()*3/4)

	// Display window
	win.ShowAll()

	return win, err
}

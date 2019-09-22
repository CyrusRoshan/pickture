package ui

import (
	"github.com/CyrusRoshan/pickture/files"
	"github.com/gobuffalo/packr"
	"github.com/gotk3/gotk3/gdk"
)

func PixbufFromFile(path string, box *packr.Box) (*gdk.Pixbuf, error) {

	//     loader = gdk_pixbuf_loader_new ();
	//     gdk_pixbuf_loader_write (loader, buffer, length, NULL);
	//     pixbuf = gdk_pixbuf_loader_get_pixbuf (loader);

	loader, _ := gdk.PixbufLoaderNew()

	b, err := files.LoadFile(path, box)
	if err != nil {
		return nil, err
	}
	return loader.WriteAndReturnPixbuf(b)
}

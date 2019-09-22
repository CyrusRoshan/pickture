package files

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/gotk3/gotk3/gdk"
)

type File struct {
	Info os.FileInfo
	Path string
}

func ListFiles(folderPath string) ([]File, error) {
	files := []File{}
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// don't add directory information
		if info.IsDir() {
			return nil
		}

		if !isValidExtension(path) {
			return nil
		}

		files = append(files, File{
			Info: info,
			Path: path,
		})
		return nil
	})
	return files, err
}

func isValidExtension(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	if extension == ".png" ||
		extension == ".jpg" ||
		extension == ".jpeg" {
		return true
	}
	return false
}

func LoadFile(path string, box *packr.Box) ([]byte, error) {
	var file http.File
	var err error

	if box == nil {
		file, err = os.Open(path)
	} else {
		file, err = box.Open(path)
	}
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func LoadImage(path string, box *packr.Box) (image.Image, error) {
	var file http.File
	var err error

	if box == nil {
		file, err = os.Open(path)
	} else {
		file, err = box.Open(path)
	}

	if err != nil {
		return nil, err
	}

	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func PixbufFromFile(path string, box *packr.Box) (*gdk.Pixbuf, error) {

	//     loader = gdk_pixbuf_loader_new ();
	//     gdk_pixbuf_loader_write (loader, buffer, length, NULL);
	//     pixbuf = gdk_pixbuf_loader_get_pixbuf (loader);

	loader, _ := gdk.PixbufLoaderNew()

	b, err := LoadFile(path, box)
	if err != nil {
		return nil, err
	}
	return loader.WriteAndReturnPixbuf(b)
}

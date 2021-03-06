package files

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/gotk3/gotk3/gdk"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
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

func LoadFile(path string, box *packr.Box) (http.File, error) {
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
	return file, nil
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
	// Create loader
	loader, err := gdk.PixbufLoaderNew()
	if err != nil {
		return nil, err
	}
	defer loader.Close()

	// Load file
	file, err := LoadFile(path, box)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Copy from file to loader
	_, err = io.Copy(loader, file)
	if err != nil {
		return nil, err
	}

	// Close loader
	err = loader.Close()
	if err != nil {
		return nil, err
	}

	// Get pixbuf from loader
	pixbuf, err := loader.GetPixbuf()
	if err != nil {
		return nil, err
	}

	return pixbuf, nil
}

func ExifDataFromFile(path string) *exif.Exif {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		x = nil
	}
	return x
}

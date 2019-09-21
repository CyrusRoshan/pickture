package files

import (
	"image"
	_ "image/jpeg"
	_ "image/png"

	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
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

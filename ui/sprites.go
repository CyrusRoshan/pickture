package ui

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"

	"github.com/faiface/pixel"
	"github.com/gobuffalo/packr"
)

func SpriteFromFile(path string, box *packr.Box) (*pixel.Sprite, error) {
	pic, err := loadPicture(path, box)
	if err != nil {
		return nil, err
	}
	return pixel.NewSprite(pic, pic.Bounds()), nil
}

func loadPicture(path string, box *packr.Box) (pixel.Picture, error) {
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
	return pixel.PictureDataFromImage(img), nil
}

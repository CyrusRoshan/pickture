package ui

import (
	"github.com/CyrusRoshan/pickture/files"
	"github.com/faiface/pixel"
	"github.com/gobuffalo/packr"
)

func PictureFromFile(path string, box *packr.Box) (*pixel.PictureData, error) {
	img, err := files.LoadImage(path, box)
	if err != nil {
		return nil, err
	}
	pic := pixel.PictureDataFromImage(img)
	return pic, nil
}

func SpriteFromFile(path string, box *packr.Box) (*pixel.Sprite, error) {
	pic, err := PictureFromFile(path, box)
	if err != nil {
		return nil, err
	}
	return pixel.NewSprite(pic, pic.Bounds()), nil
}

package ui

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

func SpriteFromFile(path string) (*pixel.Sprite, error) {
	pic, err := loadPicture(path)
	if err != nil {
		return nil, err
	}
	return pixel.NewSprite(pic, pic.Bounds()), nil
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
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

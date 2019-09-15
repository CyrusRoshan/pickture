package ui

import (
	"path/filepath"

	"github.com/CyrusRoshan/pickture/utils"
	"github.com/faiface/pixel"
)

var (
	QButton *pixel.Sprite
	WButton *pixel.Sprite
	EButton *pixel.Sprite
	AButton *pixel.Sprite
	SButton *pixel.Sprite
	DButton *pixel.Sprite
)

func init() {
	var getButtonFromAssets = func(filename string) *pixel.Sprite {
		absFilePath, err := filepath.Abs("./assets/" + filename + "_button.png")
		utils.PanicIfErr(err)

		sprite, err := SpriteFromFile(absFilePath)
		utils.PanicIfErr(err)
		return sprite
	}

	QButton = getButtonFromAssets("q")
	WButton = getButtonFromAssets("w")
	EButton = getButtonFromAssets("e")
	AButton = getButtonFromAssets("a")
	SButton = getButtonFromAssets("s")
	DButton = getButtonFromAssets("d")
}

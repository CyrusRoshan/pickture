package ui

import (
	"fmt"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel/pixelgl"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/input"
	"github.com/CyrusRoshan/pickture/utils"
)

var stateInfo = struct {
	files []files.File
}{}

func RenderCurrentState(win *pixelgl.Window) {
	// Get changes to previous state
	input.CalculateKeyPressChanges(win)

	// Apply changes to previous state
	// (none yet)

	// Draw current state
	win.Clear(colornames.Black)
	DrawButtons(win)
	DrawImageCount(win, len(stateInfo.files))
}

func SetBindings() {
	input.AddKeyPressFunctions(input.KeyPressOptions{
		Next: getNewInfo,
		A:    A,
		D:    D,
	})
}

func getNewInfo() {
	// Get all files
	fmt.Println("GET NEW INFO")
	allFiles, err := files.ListFiles()
	utils.PanicIfErr(err)
	stateInfo.files = allFiles
}

func A() {

}

func D() {

}

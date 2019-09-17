package ui

import (
	"fmt"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel/pixelgl"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/input"
	"github.com/CyrusRoshan/pickture/utils"
)

const (
	inputPath   = "./ignore/testinput"
	aOutputPath = "./ignore/testoutput/a"
	dOutputPath = "./ignore/testoutput/d"
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
	DrawBackgroundImage(win, stateInfo.files[0])
	DrawButtons(win)
	DrawImageCount(win, len(stateInfo.files))
}

func SetBindings() {
	eventChannel := input.GetKeyPressEvents()

	go func() {
		for {
			inputEvent := <-eventChannel
			switch inputEvent {
			case input.NextEvent:
				fmt.Println("Next event!")
				getNewInfo()
			case input.APressEvent:
				fmt.Println("A pressed!")
			case input.DPressEvent:
				fmt.Println("D pressed!")
			default:
				fmt.Printf("Unhandled event: %s", inputEvent)
			}
		}
	}()
}

func getNewInfo() {
	// Get all files
	allFiles, err := files.ListFiles(inputPath)
	utils.PanicIfErr(err)
	stateInfo.files = allFiles
}

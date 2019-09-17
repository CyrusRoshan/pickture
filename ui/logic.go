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
				MoveFile(inputPath, aOutputPath)
			case input.DPressEvent:
				fmt.Println("D pressed!")
				MoveFile(inputPath, dOutputPath)
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

type FileAction string

const Move = FileAction("mv")
const Copy = FileAction("cp")
const Delete = FileAction("rm")

type Change struct {
	FileAction FileAction
	Arg1       string
	Arg2       string
}

func ReverseChange(c Change) Change {
	var reverse Change
	switch c.FileAction {
	case Move:
		reverse.FileAction = Move
		reverse.Arg1 = c.Arg2
		reverse.Arg2 = c.Arg1
	case Copy:
		reverse.FileAction = Delete
		reverse.Arg1 = c.Arg2
	}
	return reverse
}

var changes = []Change{}

func ExecuteChange(c Change) {
	// Add change execution stuff here
}

func MoveFile(originalPath string, newPath string) func() {
	return func() {
		c := Change{
			FileAction: Move,
			Arg1:       originalPath,
			Arg2:       newPath,
		}
		changes = append(changes, c)
		ExecuteChange(c)
	}
}

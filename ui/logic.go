package ui

import (
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
	input.AddKeyPressFunctions(input.KeyPressOptions{
		Next: getNewInfo,
		A:    MoveFile(inputPath, aOutputPath),
		D:    MoveFile(inputPath, dOutputPath),
	})
}

func getNewInfo() {
	// Get all files
	allFiles, err := files.ListFiles(inputPath)
	utils.PanicIfErr(err)
	stateInfo.files = allFiles
}

type Action string

const Move = Action("mv")
const Copy = Action("cp")
const Delete = Action("rm")

type Change struct {
	Action Action
	Arg1   string
	Arg2   string
}

func ReverseChange(c Change) Change {
	var reverse Change
	switch c.Action {
	case Move:
		reverse.Action = Move
		reverse.Arg1 = c.Arg2
		reverse.Arg2 = c.Arg1
	case Copy:
		reverse.Action = Delete
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
			Action: Move,
			Arg1:   originalPath,
			Arg2:   newPath,
		}
		changes = append(changes, c)
		ExecuteChange(c)
	}
}

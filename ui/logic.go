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

func currentFile() *files.File {
	if len(stateInfo.files) == 0 {
		return nil
	}
	return &stateInfo.files[0]
}

func RenderCurrentState(win *pixelgl.Window) {
	// Get changes to previous state
	input.CalculateKeyPressChanges(win)

	// Apply changes to previous state
	// (none yet)

	// Draw current state
	win.Clear(colornames.Black)
	// Don't draw background if there is no image
	currFile := currentFile()
	if currFile != nil {
		DrawBackgroundImage(win, *currFile)
	}
	DrawButtons(win)
	DrawImageCount(win, len(stateInfo.files))
}

func SetBindings() {
	eventChannel := input.GetKeyPressEvents()

	// All changes that have happened
	previousChanges := []files.Change{}

	// The current change
	currentChange := files.Change{}
	addPathToCopyTo := func(pathPrefix string) {
		fullPath := pathPrefix + "/" + currentFile().Info.Name()
		currentChange.NewPaths = append(currentChange.NewPaths, fullPath)
	}

	go func() {
		for {
			inputEvent := <-eventChannel

			if inputEvent != input.UndoEvent {
				// Ignore keypresses if we have no files left in the folder
				if currentFile() == nil {
					continue
				}
			}

			switch inputEvent {
			case input.NextEvent:
				fmt.Println("Next event!")
				// Set change src to current file
				currentChange.OriginalPath = currentFile().Path

				// Get execute change commands
				cmds := files.GetChangeCommands(currentChange)
				files.ExecuteChangeCommands(cmds)

				// Add change to list of prev changes, reset curr change
				previousChanges = append(previousChanges, currentChange)
				currentChange = files.Change{}

				// Update the new info
				getNewInfo()
			case input.UndoEvent:
				fmt.Println("UNDO EVENT!!!")
				// If current change isn't empty, just reset it and exit
				if len(currentChange.NewPaths) != 0 {
					currentChange = files.Change{}
					break
				}

				// If empty, exit
				if len(previousChanges) == 0 {
					break
				}

				// If normal, get prev change, pop it, and undo it
				changeToUndo := previousChanges[len(previousChanges)-1]
				previousChanges = previousChanges[:len(previousChanges)-1]

				cmds := files.GetReverseChangeCommands(changeToUndo)
				files.ExecuteChangeCommands(cmds)

				// Update the new info
				getNewInfo()

			case input.APressEvent:
				fmt.Println("A pressed!")
				addPathToCopyTo(aOutputPath)

			case input.DPressEvent:
				fmt.Println("D pressed!")
				addPathToCopyTo(dOutputPath)

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

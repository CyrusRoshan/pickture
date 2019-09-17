package logic

import (
	"fmt"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/input"
)

const (
	inputPath   = "./ignore/testinput"
	aOutputPath = "./ignore/testoutput/a"
	dOutputPath = "./ignore/testoutput/d"
)

func Init() {
	getNewState() // Update state when starting

	eventChannel := input.GetKeyPressEvents()

	// All changes that have happened
	previousChanges := []files.Change{}

	// The current change
	currentChange := files.Change{}
	addPathToCopyTo := func(pathPrefix string) {
		fullPath := pathPrefix + "/" + CurrentFile().Info.Name()
		currentChange.NewPaths = append(currentChange.NewPaths, fullPath)
	}

	go func() {
		for {
			inputEvent := <-eventChannel

			if inputEvent != input.UndoEvent {
				// Ignore keypresses if we have no files left in the folder
				if CurrentFile() == nil {
					continue
				}
			}

			switch inputEvent {
			case input.NextEvent:
				fmt.Println("Next event!")
				// Set change src to current file
				currentChange.OriginalPath = CurrentFile().Path

				// Get execute change commands
				cmds := files.GetChangeCommands(currentChange)
				files.ExecuteChangeCommands(cmds)

				// Add change to list of prev changes, reset curr change
				previousChanges = append(previousChanges, currentChange)
				currentChange = files.Change{}

				// Update the state
				getNewState()
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

				// Update the state
				getNewState()

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

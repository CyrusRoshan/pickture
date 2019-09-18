package logic

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/input"
)

type InitProperties struct {
	InputPath           string
	AOutputPath         string
	DOutputPath         string
	DisableUniqueSuffix bool
}

var props InitProperties

func Init(p InitProperties) {
	props = p
	getNewState() // Update state when starting

	eventChannel := input.GetKeyPressEvents()

	// All changes that have happened
	previousChanges := []files.Change{}

	// The current change
	currentChange := files.Change{}
	addPathToCopyTo := func(pathPrefix string) {
		var outputName string
		inputName := CurrentFile().Info.Name()

		if props.DisableUniqueSuffix {
			outputName = inputName
		} else {
			// Get file name
			ext := filepath.Ext(inputName)
			prefix := strings.TrimSuffix(inputName, ext)

			// Make unique output file name, given input file name
			// Note: whether this goes to folder a or d, the output
			// file name will be the same, to allow you to reference
			// the same file across output folders
			outputName = fmt.Sprintf("%s.%s%s", prefix, stateInfo.fileUUID, ext)
		}

		fullPath := pathPrefix + "/" + outputName
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
				addPathToCopyTo(props.AOutputPath)

			case input.DPressEvent:
				fmt.Println("D pressed!")
				addPathToCopyTo(props.DOutputPath)

			default:
				fmt.Printf("Unhandled event: %s", inputEvent)
			}
		}
	}()
}

package logic

import (
	"fmt"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/input"
)

func handleInputEvents(inputEvents chan input.InputEvent) {
	for {
		inputEvent := <-inputEvents

		// Ignore keypresses if we have no files left in the folder
		if inputEvent != input.UndoEvent {
			if State.GetCurrentFile() == nil {
				continue
			}
		}
		fmt.Printf("New Event: %s!\n", inputEvent)

		switch inputEvent {
		case input.NextEvent:
			onNext()
		case input.UndoEvent:
			onUndo()
		case input.APressEvent:
			State.addOutputPathToCurrentChange(props.AOutputPath)
		case input.SPressEvent:
			State.addOutputPathToCurrentChange(props.SOutputPath)
		case input.DPressEvent:
			State.addOutputPathToCurrentChange(props.DOutputPath)
		case input.QPressEvent:
			State.addOutputPathToCurrentChange(props.QOutputPath)
		case input.WPressEvent:
			State.addOutputPathToCurrentChange(props.WOutputPath)
		case input.EPressEvent:
			State.addOutputPathToCurrentChange(props.EOutputPath)
		default:
			fmt.Printf("Unhandled event: %s", inputEvent)
		}
	}
}

func onNext() {
	// Set change src to current file
	State.currentChange.OriginalPath = State.GetCurrentFile().Path

	// Get execute change commands
	cmds := files.GetChangeCommands(State.currentChange)
	files.ExecuteChangeCommands(cmds)

	// Add change to list of prev changes, reset curr change
	State.previousChanges = append(State.previousChanges, State.currentChange)
	State.currentChange = files.Change{}

	// Update the state
	State.nextFile()
	State.update()
}

func onUndo() {
	// If current change isn't empty, just reset it and exit
	if len(State.currentChange.NewPaths) != 0 {
		State.currentChange = files.Change{}
		return
	}
	// If empty, exit
	if len(State.previousChanges) == 0 {
		return
	}

	// If normal, get prev change, pop it, and undo it
	changeToUndo := State.previousChanges[len(State.previousChanges)-1]
	State.previousChanges = State.previousChanges[:len(State.previousChanges)-1]

	cmds := files.GetReverseChangeCommands(changeToUndo)
	files.ExecuteChangeCommands(cmds)

	// Update the state
	State.prevFile()
	State.update()

}

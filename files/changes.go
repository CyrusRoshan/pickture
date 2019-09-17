package files

import (
	"os"

	"github.com/CyrusRoshan/pickture/utils"
)

type fileAction string

const move = fileAction("mv")
const copy = fileAction("cp")
const delete = fileAction("rm")

type Change struct {
	OriginalPath string   // the original path of the file
	NewPaths     []string // the new [...copied, moved] file paths
}

type ChangeCommand struct {
	Action fileAction
	Path1  string
	Path2  string
}

func GetChangeCommands(c Change) []ChangeCommand {
	commands := make([]ChangeCommand, len(c.NewPaths))

	var i int
	for i = 0; i < len(c.NewPaths)-1; i++ {
		commands[i] = ChangeCommand{
			Action: copy,
			Path1:  c.OriginalPath,
			Path2:  c.NewPaths[i],
		}
	}
	commands[i] = ChangeCommand{
		Action: move,
		Path1:  c.OriginalPath,
		Path2:  c.NewPaths[i],
	}

	return commands
}

func GetReverseChangeCommands(c Change) []ChangeCommand {
	commands := make([]ChangeCommand, len(c.NewPaths))

	var i int
	for i = 0; i < len(c.NewPaths)-1; i++ {
		commands[i] = ChangeCommand{
			Action: delete,
			Path1:  c.NewPaths[i],
		}
	}
	commands[i] = ChangeCommand{
		Action: move,
		Path1:  c.NewPaths[i],
		Path2:  c.OriginalPath,
	}

	return commands
}

var changes = []Change{}

func ExecuteChangeCommands(ccs []ChangeCommand) {
	for _, change := range ccs {
		switch change.Action {
		case move:
			err := os.Rename(change.Path1, change.Path2)
			utils.PanicIfErr(err)
		case copy:
			err := os.Link(change.Path1, change.Path2)
			utils.PanicIfErr(err)
		case delete:
			err := os.Remove(change.Path1)
			utils.PanicIfErr(err)
		}
	}
}

package files

import (
	"os"
	"path"

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

func ExecuteChangeCommands(ccs []ChangeCommand) {
	createParentFolder := func(filePath string) {
		// Get parent folder
		dir := path.Dir(filePath)

		// Create only if it doesn't exist
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, os.ModePerm)
			utils.PanicIfErr(err)
		}
	}

	for _, change := range ccs {
		switch change.Action {
		case move:
			createParentFolder(change.Path2)
			err := os.Rename(change.Path1, change.Path2)
			utils.PanicIfErr(err)
		case copy:
			createParentFolder(change.Path2)
			err := os.Link(change.Path1, change.Path2)
			utils.PanicIfErr(err)
		case delete:
			err := os.Remove(change.Path1)
			utils.PanicIfErr(err)
		}
	}
}

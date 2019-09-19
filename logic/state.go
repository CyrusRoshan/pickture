package logic

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/google/uuid"
)

var State = state{}

type state struct {
	fileIndex int
	files     []files.File

	// UUID for moved files (to avoid conflicts)
	fileUUID string

	// All changes that have happened
	previousChanges []files.Change
	// The current change
	currentChange files.Change
}

// GetImageCount used for displaying image count
func (s *state) GetImageCount() int {
	return len(s.files)
}

// GetCurrentFile used for displaying current file name
func (s *state) GetCurrentFile() *files.File {
	if s.fileIndex > len(s.files)-1 {
		return nil
	}
	return &s.files[s.fileIndex]
}

func (s *state) initialize() {
	// Get all files
	allFiles, err := files.ListFiles(props.InputPath)
	utils.PanicIfErr(err)
	s.files = allFiles

	if !props.DisableUniqueSuffix {
		s.newFileUUID()
	}
}

func (s *state) update() {
	if !props.DisableUniqueSuffix {
		s.newFileUUID()
	}
}

func (s *state) nextFile() {
	// Allow traversing 1 past the bounds of the array,
	// which is where we return nil values from currentFile()
	if s.fileIndex < len(s.files) {
		s.fileIndex++
	}
}

func (s *state) prevFile() {
	if s.fileIndex > 0 {
		s.fileIndex--
	}
}

func (s *state) newFileUUID() {
	id := uuid.New()
	cleanId := strings.ReplaceAll(id.String(), "-", "")
	s.fileUUID = cleanId[:15] // how much do we need?
}

func (s *state) addOutputPathToCurrentChange(path string) {
	var outputName string
	inputName := s.GetCurrentFile().Info.Name()

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
		outputName = fmt.Sprintf("%s.%s%s", prefix, s.fileUUID, ext)
	}

	fullPath := path + "/" + outputName
	s.currentChange.NewPaths = append(s.currentChange.NewPaths, fullPath)
}

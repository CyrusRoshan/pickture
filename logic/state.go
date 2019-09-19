package logic

import (
	"strings"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/google/uuid"
)

var stateInfo = struct {
	fileIndex int
	files     []files.File

	fileUUID string
}{}

func getInitialState() {
	// Get all files
	allFiles, err := files.ListFiles(props.InputPath)
	utils.PanicIfErr(err)
	stateInfo.files = allFiles

	if !props.DisableUniqueSuffix {
		stateInfo.fileUUID = generateFileUUID()
	}
}

func getNewState() {
	if !props.DisableUniqueSuffix {
		stateInfo.fileUUID = generateFileUUID()
	}
}

func GetImageCount() int {
	return len(stateInfo.files)
}

func GetCurrentFile() *files.File {
	if stateInfo.fileIndex > len(stateInfo.files)-1 {
		return nil
	}
	return &stateInfo.files[stateInfo.fileIndex]
}

func nextFile() {
	// Allow traversing 1 past the bounds of the array,
	// which is where we return nil values from currentFile()
	if stateInfo.fileIndex < len(stateInfo.files) {
		stateInfo.fileIndex++
	}
}

func prevFile() {
	if stateInfo.fileIndex > 0 {
		stateInfo.fileIndex--
	}
}

func generateFileUUID() string {
	id := uuid.New()
	cleanId := strings.ReplaceAll(id.String(), "-", "")
	return cleanId[:15] // how much do we need?
}

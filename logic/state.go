package logic

import (
	"strings"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/google/uuid"
)

var stateInfo = struct {
	files    []files.File
	fileUUID string
}{}

func CurrentFile() *files.File {
	if len(stateInfo.files) == 0 {
		return nil
	}
	return &stateInfo.files[0]
}

func GetImageCount() int {
	return len(stateInfo.files)
}

func getNewState() {
	// Get all files
	allFiles, err := files.ListFiles(props.InputPath)
	utils.PanicIfErr(err)
	stateInfo.files = allFiles

	if !props.DisableUniqueSuffix {
		stateInfo.fileUUID = generateFileUUID()
	}
}

func generateFileUUID() string {
	id := uuid.New()
	cleanId := strings.ReplaceAll(id.String(), "-", "")
	return cleanId[:15] // how much do we need?
}

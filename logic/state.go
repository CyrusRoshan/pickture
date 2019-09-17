package logic

import (
	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/utils"
)

var stateInfo = struct {
	files []files.File
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

func getNewState(inputPath string) {
	// Get all files
	allFiles, err := files.ListFiles(inputPath)
	utils.PanicIfErr(err)
	stateInfo.files = allFiles
}

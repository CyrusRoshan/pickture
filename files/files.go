package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const inputFolder = "./ignore/testinput"
const outputFolder = "./ignore/testoutput"

type File struct {
	Info os.FileInfo
	Path string
}

func ListFiles() ([]File, error) {
	files := []File{}
	err := filepath.Walk(inputFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// don't add directory information
		if info.IsDir() {
			return nil
		}

		if !isValidExtension(path) {
			return nil
		}

		files = append(files, File{
			Info: info,
			Path: path,
		})
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	return files, err
}

func isValidExtension(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	if extension == ".png" ||
		extension == ".jpg" ||
		extension == ".jpeg" ||
		extension == ".raw" ||
		extension == ".nef" {
		return true
	}
	return false
}

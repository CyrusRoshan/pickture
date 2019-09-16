package files

import (
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Info os.FileInfo
	Path string
}

func ListFiles(folderPath string) ([]File, error) {
	files := []File{}
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
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

package logic

import "github.com/CyrusRoshan/pickture/files"

type fileList struct {
	fileIndex int
	files     []files.File
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

package logic

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/rwcarlsen/goexif/exif"

	"github.com/CyrusRoshan/pickture/cache"

	"github.com/CyrusRoshan/pickture/files"
	"github.com/CyrusRoshan/pickture/utils"
	"github.com/google/uuid"
)

var State = state{}

type state struct {
	// UUID for moved files (to avoid conflicts)
	fileUUID string

	// To hold change info (for undo functionality)
	previousChanges []files.Change
	currentChange   files.Change

	// Simplify getting files
	fileList
	fileDataCache *cache.FileArrayCache
}

type ImageData struct {
	Pixbuf   *gdk.Pixbuf
	ExifData *exif.Exif
}

func (s *state) initialize() {
	// Get all files
	allFiles, err := files.ListFiles(props.InputPath)
	utils.PanicIfErr(err)
	s.files = allFiles

	// Set up the cache
	paths := make([]string, len(s.files))
	for i, file := range s.files {
		paths[i] = file.Path
	}
	s.fileDataCache = cache.NewFileArrayCache(cache.FileArrayCacheProps{
		Paths: paths,
		LoadFunc: func(path string) interface{} {
			pixbuf, err := files.PixbufFromFile(path, nil)
			utils.PanicIfErr(err, "Error loading image for cache")

			exifData := files.ExifDataFromFile(path)

			return ImageData{
				Pixbuf:   pixbuf,
				ExifData: exifData,
			}
		},
		PreloadCount: 5,
	})

	if !props.DisableUniqueSuffix {
		s.resetFileUUID()
	}
}

func (s *state) update() {
	if !props.DisableUniqueSuffix {
		s.resetFileUUID()
	}
}

// GetImageCount used for displaying image count
func (s *state) GetImageCount() int {
	return len(s.files)
}

// GetCurrentImageIndex used for displaying image count
func (s *state) GetCurrentImageIndex() int {
	return s.fileIndex
}

// GetCurrentFile used for displaying current file name
func (s *state) GetCurrentFile() *files.File {
	if s.fileIndex > len(s.files)-1 || s.fileIndex < 0 {
		return nil
	}
	return &s.files[s.fileIndex]
}

// GetCurrentFile used for displaying current file name
func (s *state) GetCurrentImage() *ImageData {
	if s.fileIndex > len(s.files)-1 || s.fileIndex < 0 {
		return nil
	}
	fileData := s.fileDataCache.Get(s.fileIndex).(ImageData)
	return &fileData
}

func (s *state) resetFileUUID() {
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

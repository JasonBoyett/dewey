package test

import (
	"errors"
	"os"
	"path/filepath"
)

// the setup function creates a temporary file structure for testing
func setup() (os.File, error) {
	baseDir, err := os.Getwd()
	rootPath := filepath.Join(baseDir, "test", "test_dir")
	if err != nil {

		return os.File{}, err
	}
	rootFile, err := os.Open(rootPath)
	if err != nil {
		err = errorCleanup(err, rootPath)
		if errors.Is(err, CleanupError{}) {
			return os.File{}, err
		}
		return os.File{}, SetupError{Err: err}
	}
	return *rootFile, err
}

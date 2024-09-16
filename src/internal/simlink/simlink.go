package simlink

import (
	"os"
	"path/filepath"

	search "github.com/jasonboyett/dewey/src/internal/search"
)

func CreateSimLink(src string, dest string) error {
	destInfo, err := os.Stat(dest)
	if err != nil {
		return SimLinkCreationError{Err: err}
	}
	if destInfo.IsDir() {
		dest = filepath.Join(dest, filepath.Base(src))
	}
	if err := os.Symlink(src, dest); err != nil {
		return SimLinkCreationError{Err: err}
	}
	return nil
}

func GroupSimLinks(
	fileExtensions []string,
	startPath string,
	dst string,
	excludeFiles []string,
) error {
	// Search for files with the given extensions
	paths := []string{}
	for _, ext := range fileExtensions {
		results, err := search.Search(startPath, ext)
		if err != nil {
			return err
		}
		paths = append(paths, results...)
	}
	// Create symlinks for each file found
	for _, path := range paths {
		if contains(excludeFiles, path) {
			continue
		}
		if err := CreateSimLink(path, dst); err != nil {
			return err
		}
	}
	return nil
}

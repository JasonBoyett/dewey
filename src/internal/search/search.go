package search

import (
	"os"
	"path/filepath"
	"strings"
)

// Search the file system for files with the given extension
// param start: the directory to start the search
// param extension: the extension to search for
// return: a slice of file paths with the given extension
func Search(start string, extension string) ([]string, error) {
	var files []string
	err := filepath.Walk(start, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == extension {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, FileSearchError{Err: err}
	}

	return excludeHidden(files), nil
}

func excludeHidden(unfiltered []string) []string {
	var filtered []string
	for _, path := range unfiltered {
		if !isHidden(path) {
			filtered = append(filtered, path)
		}
	}
	return filtered
}

func isHidden(path string) bool {
	for _, part := range strings.Split(path, string(os.PathSeparator)) {
		if strings.HasPrefix(part, ".") {
			return true
		}
	}
	return false
}

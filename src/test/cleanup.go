package test

import (
	"bufio"
	"os"
	"path/filepath"
)

func errorCleanup(trigger error, rootPath string) error {
	fileText := []string{
		"The rest of this directory should be empty.\n",
		"If there are any other files or directories in this directory the previous\n",
		"test failed to clean up.\n",
		"Do not manually create any files or directories in this directory.\n",
	}

	rootDir := filepath.Dir(rootPath)
	err := os.RemoveAll(rootDir)
	if err != nil {
		return CleanupError{Err: err}
	}

	err = os.MkdirAll(rootDir, 0755)
	if err != nil {
		return CleanupError{Err: err}
	}

	fileName := filepath.Join(rootPath, "root.txt")
	rootFile, err := os.Create(fileName)
	defer rootFile.Close()
	if err != nil {
		return CleanupError{Err: err}
	}

	writer := bufio.NewWriter(rootFile)
	defer writer.Flush()
	for _, line := range fileText {
		_, err = writer.WriteString(line)
		if err != nil {
			return CleanupError{Err: err}
		}
	}

	return trigger
}

func cleanup(trigger error) error {
	// TODO: implement cleanup function
	return trigger
}

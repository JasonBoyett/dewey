package test

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func errorCleanup(trigger error, rootPath string) error {
	// fmt.Println("errorCleanup")
	// err := cleanup(rootPath)
	// if err != nil {
	// 	if errors.Is(err, CleanupError{}) {
	// 		return err
	// 	} else {
	// 		return CleanupError{Err: err}
	// 	}
	// }
	return trigger
}

func cleanup(rootPath string) error {
	fileText := []string{
		"The rest of this directory should be empty.",
		"If there are any other files or directories in this directory the previous",
		"test failed to clean up.",
		"Do not manually create any files or directories in this directory.",
	}
	rootPath, err := backtrackTestDir("test_dir", rootPath)
	if err != nil {
		return CleanupError{Err: err}
	}

	err = os.RemoveAll(rootPath)
	if err != nil {
		return CleanupError{Err: err}
	}

	err = os.MkdirAll(rootPath, 0755)
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
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return CleanupError{Err: err}
		}
	}
	return nil
}

func backtrackTestDir(tgt, dir string) (string, error) {
	for {
		if !strings.Contains(dir, tgt) {
			return "", errors.New(fmt.Sprintf("%s not found in %s", tgt, dir))
		}
		if filepath.Base(dir) == tgt {
			return dir, nil
		}
		dir = filepath.Dir(dir)
		if dir == "/" {
			return "", errors.New(fmt.Sprintf("%s not found", tgt))
		}
	}
}

package simlink

import "os"

func contains[T comparable](slice []T, item T) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func IsSimLink(path string) (bool, error) {
	file, err := os.Lstat(path)
	if err != nil {
		return false, err
	}
	return file.Mode()&os.ModeSymlink == os.ModeSymlink, nil
}

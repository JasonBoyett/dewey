package shelf

import (
	"os"
	"path/filepath"

	simlink "github.com/jasonboyett/dewey/src/internal/simlink"
)

type Shelf struct {
	Name      string   `json:"name"`
	FileTypes []string `json:"fileTypes"`
}

func (s *Shelf) Populate(lib *Library) error {
	dst := filepath.Join(lib.Path, s.Name)
	currentSymlinks, err := s.getCurrentSymlinks(lib)
	if err != nil {
		return err
	}
	if err := simlink.GroupSimLinks(
		s.FileTypes,
		lib.SearchBase,
		dst,
		currentSymlinks,
	); err != nil {
		return err
	}
	return nil
}

func (s *Shelf) getCurrentSymlinks(lib *Library) ([]string, error) {
	var symlinks []string
	dst := filepath.Join(lib.Path, s.Name)
	content, err := os.ReadDir(dst)
	if err != nil {
		return nil, err
	}
	for _, entry := range content {
		if entry.IsDir() {
			continue
		}
		path := filepath.Join(dst, entry.Name())
		isSymlink, err := simlink.IsSimLink(path)
		if err != nil {
			return nil, err
		}
		if isSymlink {
			symlinks = append(symlinks, path)
		}
	}
	return symlinks, nil
}

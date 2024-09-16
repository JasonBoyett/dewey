package shelf

import (
	"encoding/json"
	"os"
)

type Library struct {
	Shelfs []Shelf `json:"shelfs"`
	Path   string  `json:"path"`
	// SearchBase is the directory where the library starts searching for files
	// to create symlinks for.
	SearchBase string `json:"base"`
}

func (l *Library) AddShelf(s Shelf) {
	l.Shelfs = append(l.Shelfs, s)
}

func (l *Library) Update(name string) {
	for i, s := range l.Shelfs {
		if s.Name == name {
			l.Shelfs[i].Populate(l)
		}
	}
}

func (l *Library) Save(name string) error {
	libPath, err := findDeffenitionPath()
	if err != nil {
		return LibrarySaveError{Err: err}
	}
	jsonBytes, err := json.Marshal(l)
	if err != nil {
		return err
	}
	if err = os.WriteFile(libPath, jsonBytes, 0644); err != nil {
		return LibrarySaveError{Err: err}
	}
	return nil
}

func LoadLibrary() (Library, LibraryError) {
	libPath, err := findDeffenitionPath()
	if err != nil {
		return Library{}, LibraryLoadError{Err: err}
	}
	jsonBytes, err := readLibFile(libPath)
	if err != nil {
		return Library{}, LibraryLoadError{Err: err}
	}
	return decodeLibrary(jsonBytes)
}

func CreateLibrary(base string) (Library, LibraryError) {
	path, err := findDeffenitionPath()
	if err != nil {
		return Library{}, LibraryCreationError{Err: err}
	}
	lib := Library{
		SearchBase: base,
		Path:       path,
	}
	jsonBytes, err := json.Marshal(lib)
	if err != nil {
		return Library{}, LibraryCreationError{Err: err}
	}
	if err = os.WriteFile(path, jsonBytes, 0644); err != nil {
		return Library{}, LibraryCreationError{Err: err}
	}
	return lib, nil
}

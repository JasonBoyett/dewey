package shelf

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

func findDeffenitionPath() (string, error) {
	binaryPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	libPath := filepath.Dir(binaryPath)
	// the binary is in the bin directory
	// the library is in the store directory which is in the parent directory of the bin directory
	// so we need to go up one directory by calling filepath.Dir twice
	libPath = filepath.Dir(libPath)
	libPath = filepath.Join(libPath, "store.json")
	// Now we make sure the library file exists and create it if it doesn't
	_, err = os.Stat(libPath)
	if os.IsNotExist(err) {
		err = createLibraryFile(libPath)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	return libPath, nil
}

func createLibraryFile(libPath string) error {
	libDir := filepath.Dir(libPath)
	err := os.MkdirAll(libDir, 0755)
	if err != nil {
		return err
	}
	_, err = os.Create(libPath)
	if err != nil {
		return err
	}
	return nil
}

func readLibFile(libPath string) ([]byte, error) {
	jsonFile, err := os.Open(libPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func decodeLibrary(jsonBytes []byte) (Library, LibraryError) {
	var lib Library
	err := json.Unmarshal(jsonBytes, &lib)
	if err != nil {
		return Library{}, LibraryGenerationError{Err: err}
	}
	return lib, nil
}

package test

import (
	"os"
	"testing"
)

type TestInfo struct {
	RootPath string
	RootFile os.File
	// the keys of the testContent map are paths to test files and the
	// values are the expected content of the files
	PathMap        map[string][]string
	TestLibPath    string
	TestConfigPath string
	Testing        *testing.T
}

// GetFileMap returns a map of os.File objedts as keys and their expected
// content as values. The file objects are open and should be closed by the
// caller.
func (ti *TestInfo) GetFileMap() (map[os.File][]string, error) {
	result := map[os.File][]string{}
	for path, content := range ti.PathMap {
		file, err := os.Open(path)
		if err != nil {
			return map[os.File][]string{}, err
		}
		result[*file] = content
	}
	return result, nil
}

func (ti *TestInfo) Cleanup() {
	err := os.RemoveAll(ti.RootPath)
	if err != nil {
		ti.Testing.Errorf("Error cleaning up test config path: %v", err)
	}

	err = os.RemoveAll(ti.TestLibPath)
	if err != nil {
		ti.Testing.Errorf("Error cleaning up test config path: %v", err)
	}

	err = os.RemoveAll(ti.TestConfigPath)
	if err != nil {
		ti.Testing.Errorf("Error cleaning up test config path: %v", err)
	}
	ti.Testing.Log("Cleanup complete")
}

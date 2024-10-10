package test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// the setup function creates a temporary file structure for testing
//
// The test file system will be set up like this:
//
//	test_config/ <- the mock dewey configuration directory
//	test_libray/ <- the mock dewey library directory
//	test_dir/ <- the root of the mock file system
//	├── root.txt
//	├── dir1/
//	│   ├── file1.txt
//	│   └── image1.png
//	└── dir2/
//		     ├── sub_dir1/
//				 │   ├── word_doc1.docx
//			   │   ├── word_doc2.docx
//				 │   ├── some_image.jpg
//				 │   └── word_doc3.docx
//				 ├── sub_dir2/
//				 │   ├── some_code1.rb
//				 │   ├── some_code2.rb
//				 │   ├── some_code3.rb
//				 │   └── SomeCode4.java
//				 └── .hidden_dir/
//				     ├── hidden_file1.txt
//				     ├── hidden_file2.txt
//				     ├── hidden_file3.txt
//				     └── hidden_file4.txt
//
// root.txt should have the following text:
// "Hello from the Dewey test suite!"
func setup(t *testing.T) (TestInfo, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return TestInfo{}, err
	}

	configPath := filepath.Join(baseDir, "test_config")
	err = os.MkdirAll(configPath, 0755)
	if err != nil {
		return TestInfo{}, SetupError{Err: err}
	}

	libraryPath := filepath.Join(baseDir, "test_library")
	err = os.MkdirAll(libraryPath, 0755)
	if err != nil {
		return TestInfo{}, SetupError{Err: err}
	}

	rootPath := filepath.Join(baseDir, "test_dir")
	err = os.MkdirAll(rootPath, 0755)
	if err != nil {
		return TestInfo{}, SetupError{Err: err}
	}
	rootFile, err := os.Open(rootPath)
	if err != nil {
		err = errorCleanup(err, rootPath)
		if errors.Is(err, CleanupError{}) {
			return TestInfo{}, err
		}
		return TestInfo{}, SetupError{Err: err}
	}

	files, err := buildFileStructure(rootPath)
	if err != nil {
		return TestInfo{}, SetupError{Err: err}
	}

	err = fillFiles(files)
	if err != nil {
		return TestInfo{}, SetupError{Err: err}
	}

	result := TestInfo{
		RootPath:       rootFile.Name(),
		RootFile:       *rootFile,
		PathMap:        setupTestContent(files),
		TestLibPath:    libraryPath,
		TestConfigPath: configPath,
		Testing:        t,
	}
	return result, err
}

func setupTestContent(files map[os.File][]string) map[string][]string {
	result := map[string][]string{}
	for file, content := range files {
		result[file.Name()] = content
	}
	return result
}

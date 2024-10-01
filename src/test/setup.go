package test

import (
	"errors"
	"os"
	"path/filepath"
)

// the setup function creates a temporary file structure for testing
//
// The test file system will be set up like this:
//
//	test_dir/
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
func setup() (os.File, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return os.File{}, err
	}
	rootPath := filepath.Join(baseDir, "test_dir")
	rootFile, err := os.Open(rootPath)
	if err != nil {
		err = errorCleanup(err, rootPath)
		if errors.Is(err, CleanupError{}) {
			return os.File{}, err
		}
		return os.File{}, SetupError{Err: err}
	}

	files, err := buildFileStructure(rootPath)
	if err != nil {
		return os.File{}, SetupError{Err: err}
	}

	err = fillFiles(files, rootPath)
	if err != nil {
		return os.File{}, SetupError{Err: err}
	}

	return *rootFile, err
}

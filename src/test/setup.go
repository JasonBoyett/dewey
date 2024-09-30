package test

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
)

// The test file system will be set up like this:
//
// test_dir/
// ├── root.txt
// ├── dir1/
// │   ├── file1.txt
// │   └── image1.png
// └── dir2/
//	     ├── sub_dir1/
//			 │   ├── word_doc1.docx
//		   │   ├── word_doc2.docx
//			 │   ├── some_image.jpg
//			 │   └── word_doc3.docx
//			 ├── sub_dir2/
//			 │   ├── some_code1.rb
//			 │   ├── some_code2.rb
//			 │   ├── some_code3.rb
//			 │   └── some_code4.java
//			 └── .hidden_dir/
//			     ├── hidden_file1.txt
//			     ├── hidden_file2.txt
//			     ├── hidden_file3.txt
//			     └── hidden_file4.txt
// root.txt should have the following text:
// "Hello from the Dewey test suite!"

// the setup function creates a temporary file structure for testing
func setup() (os.File, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return os.File{}, err
	}
	rootPath := filepath.Join(baseDir, "test", "test_dir")
	rootFile, err := os.Open(rootPath)
	if err != nil {
		err = errorCleanup(err, rootPath)
		if errors.Is(err, CleanupError{}) {
			return os.File{}, err
		}
		return os.File{}, SetupError{Err: err}
	}
	defer cleanup(nil)
	return *rootFile, err
}

func buildFileStructure(base string) (map[string]string, error) {
	files := map[string]string{}
	// create the root directory
	err := os.MkdirAll(base, 0755)
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create the root file
	fileName := filepath.Join(base, "root.txt")
	rootFile, err := os.Create(fileName)
	defer rootFile.Close()
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}
	files["root.txt"] = "Hello from the Dewey test suite!"

	// create dir1
	dir1 := filepath.Join(base, "dir1")
	err = os.MkdirAll(dir1, 0755)
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create file1.txt
	fileName = filepath.Join(dir1, "file1.txt")
	file1, err := os.Create(fileName)
	defer file1.Close()
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create image1.png
	fileName = filepath.Join(dir1, "image1.png")
	image1, err := os.Create(fileName)
	defer image1.Close()
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create dir2
	dir2 := filepath.Join(base, "dir2")
	err = os.MkdirAll(dir2, 0755)
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create sub_dir1
	subDir1 := filepath.Join(dir2, "sub_dir1")
	err = os.MkdirAll(subDir1, 0755)
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create word_doc1.docx
	fileName = filepath.Join(subDir1, "word_doc1.docx")
	wordDoc1, err := os.Create(fileName)
	defer wordDoc1.Close()
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create word_doc2.docx
	fileName = filepath.Join(subDir1, "word_doc2.docx")
	wordDoc2, err := os.Create(fileName)
	defer wordDoc2.Close()
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create some_image.jpg
	fileName = filepath.Join(subDir1, "some_image.jpg")
	someImage, err := os.Create(fileName)
	defer someImage.Close()
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create word_doc3.docx
	fileName = filepath.Join(subDir1, "word_doc3.docx")
	wordDoc3, err := os.Create(fileName)
	defer wordDoc3.Close()
	if err != nil {
		return map[string]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	return files, nil
}

func fillFile(file *os.File, lines []string, root string) error {
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for _, line := range lines {
		_, err := writer.WriteString(line)
		if err != nil {
			return errorCleanup(
				SetupError{Err: err},
				root,
			)
		}
	}
	return nil
}

package test

import (
	"bufio"
	"os"
	"path/filepath"
	"sync"
)

//  the file structure will be set up like this:
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

func buildFileStructure(base string) (map[os.File][]string, error) {
	files := map[os.File][]string{}
	// create the root directory
	err := os.MkdirAll(base, 0755)
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create the root file
	fileName := filepath.Join(base, "root.txt")
	rootFile, err := os.Create(fileName)
	defer rootFile.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}
	files[*rootFile] = []string{"Hello from the Dewey test suite!"}

	// create dir1
	dir1 := filepath.Join(base, "dir1")
	err = os.MkdirAll(dir1, 0755)
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create file1.txt
	fileName = filepath.Join(dir1, "file1.txt")
	file1, err := os.Create(fileName)
	defer file1.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}
	files[*file1] = []string{"Hello from the Dewey test suite!"}

	// create image1.png
	fileName = filepath.Join(dir1, "image1.png")
	image1, err := os.Create(fileName)
	defer image1.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create dir2
	dir2 := filepath.Join(base, "dir2")
	err = os.MkdirAll(dir2, 0755)
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create sub_dir1
	subDir1 := filepath.Join(dir2, "sub_dir1")
	err = os.MkdirAll(subDir1, 0755)
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create word_doc1.docx
	fileName = filepath.Join(subDir1, "word_doc1.docx")
	wordDoc1, err := os.Create(fileName)
	defer wordDoc1.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create word_doc2.docx
	fileName = filepath.Join(subDir1, "word_doc2.docx")
	wordDoc2, err := os.Create(fileName)
	defer wordDoc2.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create some_image.jpg
	fileName = filepath.Join(subDir1, "some_image.jpg")
	someImage, err := os.Create(fileName)
	defer someImage.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create word_doc3.docx
	fileName = filepath.Join(subDir1, "word_doc3.docx")
	wordDoc3, err := os.Create(fileName)
	defer wordDoc3.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	// create sub_dir2
	subDir2 := filepath.Join(dir2, "sub_dir2")
	err = os.MkdirAll(subDir2, 0755)
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	fileName = filepath.Join(subDir2, "some_code1.rb")
	someCode1, err := os.Create(fileName)
	defer someCode1.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}
	code := []string{"puts 'Hello from the Dewey test suite!'"}
	files[*someCode1] = code

	fileName = filepath.Join(subDir2, "some_code2.rb")
	someCode2, err := os.Create(fileName)
	defer someCode2.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}
	code = []string{"puts 'Hello again from the Dewey test suite!'"}
	files[*someCode2] = code

	fileName = filepath.Join(subDir2, "some_code3.rb")
	someCode3, err := os.Create(fileName)
	defer someCode3.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}
	code = []string{"puts 'Hello once again from the Dewey test suite!'"}
	files[*someCode3] = code

	fileName = filepath.Join(subDir2, "SomeCode4.java")
	someCode4, err := os.Create(fileName)
	defer someCode4.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}
	code = []string{
		"Class SomeCode4 {",
		"  public static void main(String[] args) {",
		"    System.out.println(\"Hello from the Dewey test suite!\");",
		"  }",
		"}",
	}
	files[*someCode4] = code

	// create .hidden_dir
	hiddenDir := filepath.Join(dir2, ".hidden_dir")
	err = os.MkdirAll(hiddenDir, 0755)
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	fileName = filepath.Join(subDir1, "hidden_file1.txt")
	hiddenFile1, err := os.Create(fileName)
	defer hiddenFile1.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}
	fileName = filepath.Join(subDir1, "hidden_file2.txt")
	hiddenFile2, err := os.Create(fileName)
	defer hiddenFile2.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}
	fileName = filepath.Join(subDir1, "hidden_file3.txt")
	hiddenFile3, err := os.Create(fileName)
	defer hiddenFile3.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}
	fileName = filepath.Join(subDir1, "hidden_file4.txt")
	hiddenFile4, err := os.Create(fileName)
	defer hiddenFile4.Close()
	if err != nil {
		return map[os.File][]string{}, errorCleanup(SetupError{Err: err}, base)
	}

	return files, nil
}

func fillFiles(files map[os.File][]string, root string) error {
	var wg sync.WaitGroup
	errorChan := make(chan error)

	for file, lines := range files {
		wg.Add(1)
		go fillFile(
			&file,
			lines,
			root,
			errorChan,
			&wg,
		)
	}
	go func() {
		wg.Wait()
		close(errorChan)
	}()
	for err := range errorChan {
		if err != nil {
			return SetupError{Err: err}
		}
	}
	return nil
}

func fillFile(
	file *os.File,
	lines []string,
	root string,
	ch chan<- error,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			ch <- errorCleanup(SetupError{Err: err}, root)
		}
	}
}

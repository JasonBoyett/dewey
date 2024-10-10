package test

import (
	"io"
	"os"
	"strings"
	"testing"

	search "github.com/jasonboyett/dewey/src/internal/search"
	"github.com/stretchr/testify/assert"
)

func TestSearchCount(t *testing.T) {
	testInfo, err := setup(t)
	defer testInfo.Cleanup()
	defer testInfo.RootFile.Close()
	defer t.Log(testInfo.RootPath)
	if err != nil {
		t.Fatal(err)
	}
	defer testInfo.RootFile.Close()
	start := testInfo.RootFile.Name()
	results, err := search.Search(start, ".txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, len(results))

	results, err = search.Search(start, ".java")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(results))

	results, err = search.Search(start, ".rb")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(results))
}

func TestSearchContent(t *testing.T) {
	extensions := []string{".txt", ".java", ".rb"}
	testInfo, err := setup(t)
	defer testInfo.Cleanup()
	defer testInfo.RootFile.Close()
	defer t.Log(testInfo.RootPath)
	if err != nil {
		t.Fatal(err)
	}
	start := testInfo.RootFile.Name()

	for _, ext := range extensions {
		expected := func() map[string][]string {
			result := map[string][]string{}
			for path, content := range testInfo.PathMap {
				if strings.HasSuffix(path, ext) {
					result[path] = content
				}
			}
			return result
		}()

		results, err := search.Search(start, ext)
		if err != nil {
			t.Fatal(err)
		}

		for _, path := range results {
			contentFile, err := os.Open(path)
			if err != nil {
				t.Fatal(err)
			}
			defer contentFile.Close()

			contentBytes, err := io.ReadAll(contentFile)

			expectedContent := func() string {
				if len(expected[path]) == 1 {
					return strings.Join(expected[path], "")
				}
				return strings.Join(expected[path], "")
			}()

			expectedContent = strings.ReplaceAll(expectedContent, "\n", "")
			conten := strings.ReplaceAll(string(contentBytes), "\n", "")
			assert.Equal(t, expectedContent, conten)
		}
	}
}

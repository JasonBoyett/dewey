package test

import (
	"testing"
)

func TestSearch(t *testing.T) {
	startFile, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer startFile.Close()
}

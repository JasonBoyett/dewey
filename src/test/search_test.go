package test

import (
	"testing"
)

func TestSearch(t *testing.T) {
	_, err := setup()
	if err != nil {
		t.Fatal(err)
	}
}

package test

import (
	// "github.com/stretchr/testify/assert"
	"testing"
	// simlink "github.com/jasonboyett/dewey/src/internal/simlink"
)

func TestCreateSimLink(t *testing.T) {
	testInfo, err := setup(t)
	defer testInfo.Cleanup()
	if err != nil {
		t.Fatal(err)
	}
	defer testInfo.RootFile.Close()
}

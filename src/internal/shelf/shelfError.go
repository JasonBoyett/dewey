package shelf

import (
	"fmt"
)

// ShelfError is a wrapper around the standard error interface.
// It is used to differentiate between errors that occur when generating a library
// and other errors that may occur.
// This allows the caller to handle shelf generation errors gracefully while
// still being able to handle other errors in a generic way.
type ShelfError interface {
	error
	UnwrapSelfErro() error
}

type ShelfGenerationError struct {
	Err ShelfError
}

func (e ShelfGenerationError) Error() string {
	return fmt.Sprintf("error generating shelf: %v", e.Err)
}

func (e ShelfGenerationError) UnwrapSelfError() error {
	return e.Err
}

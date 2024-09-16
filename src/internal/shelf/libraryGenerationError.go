package shelf

import (
	"fmt"
)

// LibraryError is a wrapper around the standard error interface.
// It is used to differentiate between errors that occur when generating a library
// and other errors that may occur.
// This allows the caller to handle library generation errors gracefully while
// still being able to handle other errors in a generic way.
type LibraryError interface {
	error
	UnwrapLibraryError() error
}

type LibraryGenerationError struct {
	Err error
}

func (e LibraryGenerationError) Error() string {
	return fmt.Sprintf("error generating library: %v", e.Err)
}

func (e LibraryGenerationError) UnwrapLibraryError() error {
	return e.Err
}

type LibraryLoadError struct {
	Err error
}

func (e LibraryLoadError) Error() string {
	return fmt.Sprintf("error loading library: %v", e.Err)
}

func (e LibraryLoadError) UnwrapLibraryError() error {
	return e.Err
}

type LibraryCreationError struct {
	Err error
}

func (e LibraryCreationError) Error() string {
	return fmt.Sprintf("error creating library: %v", e.Err)
}

func (e LibraryCreationError) UnwrapLibraryError() error {
	return e.Err
}

type LibrarySaveError struct {
	Err error
}

func (e LibrarySaveError) Error() string {
	return fmt.Sprintf("error saving library: %v", e.Err)
}

func (e LibrarySaveError) UnwrapLibraryError() error {
	return e.Err
}

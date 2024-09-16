package search

import (
	"fmt"
)

type SearchError interface {
	error
	UnwrapSearchError() error
}

type FileSearchError struct {
	Err error
}

func (e FileSearchError) Error() string {
	return fmt.Sprintf("error in file search: %v", e.Err)
}

func (e FileSearchError) UnwrapSearchError() error {
	return e.Err
}

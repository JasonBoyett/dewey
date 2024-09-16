package simlink

import "fmt"

type SimlinkError interface {
	error
	Unwrap() error
}

type SimLinkCreationError struct {
	Err error
}

func (e SimLinkCreationError) Error() string {
	return fmt.Sprintf("SimLinkError: %v", e.Err)
}

func (e SimLinkCreationError) Unwrap() error {
	return e.Err
}

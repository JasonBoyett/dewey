package test

type SetupError struct {
	Err error
}

func (e SetupError) Error() string {
	return e.Err.Error()
}

type CleanupError struct {
	Err error
}

func (e CleanupError) Error() string {
	return e.Err.Error()
}

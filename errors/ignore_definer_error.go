package rg_errors

// When a file spec definer is ignored the GetFileSpec function should return this error so we can distinguish it from other errors.
type IgnoreDefinerError struct {
}

func (e *IgnoreDefinerError) Error() string {
	return "file spec ignore"
}

func NewIgnoreDefinerError() *IgnoreDefinerError {
	return &IgnoreDefinerError{}
}

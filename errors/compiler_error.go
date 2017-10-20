package errors

import (
	"fmt"
)

type compilerError struct {
	Message    string
	Filename   string
	LineNumber int
}

func (e *compilerError) Error() string {
	return fmt.Sprintf("ERROR %s:%d %s", e.Filename, e.LineNumber, e.Message)
}

func New(message string, filename string, lineNumber int) error {
	return &compilerError{message, filename, lineNumber}
}

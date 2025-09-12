package module

import (
	"fmt"
)

const (
	CodeNone     string = "NONE"
	CodeExpect          = "EXPECT"
	CodeNumber          = "NUMBER"
	CodeComplex         = "COMPLEX"
	CodeNotFound        = "NOT_FOUND"
)

type ParseError struct {
	Code     string
	Message  string
	LineNo   int
	ColumnNo int
}

func (err ParseError) Error() string {
	return fmt.Sprintf("Parse error at %d:%d: %s", err.LineNo, err.ColumnNo, err.Message)
}

func notFoundError(lineNo int) *ParseError {
	err := ParseError{
		Code:     CodeNotFound,
		Message:  "Module not found",
		LineNo:   lineNo,
		ColumnNo: 0,
	}

	return &err
}

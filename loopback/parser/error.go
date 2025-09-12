package parser

import (
	"fmt"
)

const (
	codeNone string = "NONE"
	codeExpect = "EXPECT"
	codeNumber = "NUMBER"
	codeComplex = "COMPLEX"
	codeNotFound = "NOT_FOUND"
)

type ParseError struct {
	code string
	message string
	lineNo int
	columnNo int
}

func (err ParseError) Error() string {
	return fmt.Sprintf("Parse error at %d:%d: %s", err.lineNo, err.columnNo, err.message)
}

func notFoundError(lineNo int) *ParseError {
	err := ParseError {
		code: codeNotFound,
		message: "Module not found",
		lineNo: lineNo,
		columnNo: 0,
	}

	return &err
}

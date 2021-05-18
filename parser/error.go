package parser

import "fmt"

type parserError struct {
	cause string
	line  int
	pos   int
}

func (err parserError) Error() string {
	return fmt.Sprintf("%s at line %d and position %d", err.cause, err.line, err.pos)
}

func newParserErrorFromToken(cause string, token Token) error {
	return parserError{
		cause: cause,
		line:  token.Line(),
		pos:   token.Begin(),
	}
}

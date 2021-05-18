package parser

import "fmt"

type parserError struct {
	cause string
	line  int
	pos   int
}

func (err parserError) Error() string {
	return fmt.Sprintf("%s at line %d and position %d", err.cause, err.line+1, err.pos+1)
}

func newParserErrorFromToken(cause string, token Token) error {
	return parserError{
		cause: cause,
		line:  token.Line(),
		pos:   token.Begin(),
	}
}

func newParserError(cause string, line, pos int) error {
	return parserError{
		cause: cause,
		line:  line,
		pos:   pos,
	}
}

package parser

import "fmt"

// type TokenQueue interface {
// 	Peek() Token
// 	Empty() bool
// 	Pop() Token
// 	Push(Token)
// }

// FIFO Stack
type tokenQueue struct {
	tokens   []Token
	lastLine int
	lastPos  int
	lastVal  string
}

func (q *tokenQueue) Peek() Token {
	return q.tokens[0]
}

func (q *tokenQueue) Empty() bool {
	return len(q.tokens) == 0
}

func (q *tokenQueue) Pop() Token {
	if len(q.tokens) == 0 {
		panic(newParserError(fmt.Sprintf("No more tokens after '%s'", q.lastVal), q.lastLine, q.lastPos))
	}
	ret := q.tokens[0]
	q.tokens = q.tokens[1:]
	q.lastLine = ret.Line()
	q.lastPos = ret.End()
	q.lastVal = ret.Content()
	return ret
}

// func (q *tokenQueue) Push(t Token) {
// 	q.tokens = append(q.tokens, t)
// }

func newTokenQueue(tokens []Token) *tokenQueue {
	return &tokenQueue{
		tokens:   tokens,
		lastLine: tokens[0].Line(),
		lastPos:  tokens[0].Begin(),
	}
}

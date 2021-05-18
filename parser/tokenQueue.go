package parser

// type TokenQueue interface {
// 	Peek() Token
// 	Empty() bool
// 	Pop() Token
// 	Push(Token)
// }

// FIFO Stack
type tokenQueue []Token

func (q *tokenQueue) Peek() Token {
	return (*q)[0]
}

func (q *tokenQueue) Empty() bool {
	return len(*q) == 0
}

func (q *tokenQueue) Pop() Token {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *tokenQueue) Push(t Token) {
	*q = append(*q, t)
}

func newTokenQueue(tokens []Token) *tokenQueue {
	ret := tokenQueue(tokens)
	return &ret
}

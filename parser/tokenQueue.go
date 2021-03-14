package parser

// FIFO Stack
type tokenQueue []string

func (q *tokenQueue) Peek() string {
	return (*q)[0]
}

func (q *tokenQueue) Empty() bool {
	return len(*q) == 0
}

func (q *tokenQueue) Pop() string {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *tokenQueue) Push(s string) {
	*q = append(*q, s)
}

func newTokenQueue(tokens []string) *tokenQueue {
	ret := tokenQueue(tokens)
	return &ret
}

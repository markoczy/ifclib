package tokenizer

// FIFO Stack
type runeQueue string

func (q *runeQueue) Peek() rune {
	return rune((*q)[0])
}

func (q *runeQueue) Empty() bool {
	return len(*q) == 0
}

func (q *runeQueue) Pop() rune {
	ret := rune((*q)[0])
	*q = (*q)[1:]
	return ret
}

func (q *runeQueue) Push(r rune) {
	*q = runeQueue(string((*q)) + string(r))
}

func newRuneQueue(tokens string) *runeQueue {
	ret := runeQueue(tokens)
	return &ret
}

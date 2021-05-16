package tokenizer

import "strings"

func isIdentifier(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return r == '_'
}

func isNumeric(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return r == '.'
}

func createToken(queue *runeQueue, begin int, validator func(rune) bool) Token {
	var sb strings.Builder
	for !queue.Empty() && validator(queue.Peek()) {
		sb.WriteRune(queue.Pop())
	}
	return &token{
		begin:   begin,
		content: sb.String(),
	}
}

func createSingleToken(queue *runeQueue, begin int) Token {
	r := queue.Pop()
	return &token{
		begin:   begin,
		content: string(r),
	}
}

func skipWhitespace(queue *runeQueue, begin int) int {
	end := begin
	for queue.Peek() == ' ' {
		queue.Pop()
		end++
	}
	return end
}

func CreateTokens(s string) []Token {
	tokens := []Token{}
	queue := runeQueue(s)
	it := 0
	for !queue.Empty() {
		r := queue.Peek()
		switch {
		case r == ' ':
			it = skipWhitespace(&queue, it)
		case isNumeric(r):
			token := createToken(&queue, it, isNumeric)
			it = token.End() + 1
			tokens = append(tokens, token)
		case isIdentifier(r):
			token := createToken(&queue, it, isIdentifier)
			it = token.End() + 1
			tokens = append(tokens, token)
		default:
			token := createSingleToken(&queue, it)
			it++
			tokens = append(tokens, token)
		}

	}
	return tokens
}

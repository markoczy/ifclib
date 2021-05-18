package tokenizer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/markoczy/ifclib/xp/names"
)

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

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func createToken(queue *runeQueue, line, begin int, validator func(rune) bool) Token {
	var sb strings.Builder
	for !queue.Empty() && validator(queue.Peek()) {
		sb.WriteRune(queue.Pop())
	}
	return &token{
		line:    line,
		begin:   begin,
		content: sb.String(),
	}
}

func createSingleToken(queue *runeQueue, line, begin int) Token {
	r := queue.Pop()
	return &token{
		line:    line,
		begin:   begin,
		content: string(r),
	}
}

func skipSpace(queue *runeQueue, begin int) int {
	end := begin
	for !queue.Empty() && isSpace(queue.Peek()) {
		queue.Pop()
		end++
	}
	return end
}

func GetEntityDefinitions(tokens []Token) ([][]Token, error) {
	return getDefinitionTokens(tokens, names.StartEntity, names.EndEntity)
}

func GetTypeDefinitions(tokens []Token) ([][]Token, error) {
	return getDefinitionTokens(tokens, names.StartType, names.EndType)
}

func GetFunctionDefinitions(tokens []Token) ([][]Token, error) {
	return getDefinitionTokens(tokens, names.StartFunction, names.EndFunction)
}

func getDefinitionTokens(tokens []Token, start, end string) ([][]Token, error) {
	ret := [][]Token{}
	capture := false
	hasEnd := false
	var cur []Token
	for _, v := range tokens {
		switch {
		case v.Content() == start:
			capture = true
			cur = []Token{}
		case v.Content() == end:
			hasEnd = true
		case hasEnd:
			if v.Content() != ";" {
				return ret, fmt.Errorf("Expected ';' after END_TYPE at token %v", v)
			}
			cur = append(cur, v)
			ret = append(ret, cur)
			capture = false
			hasEnd = false
		}
		if capture {
			cur = append(cur, v)
		}
	}
	return ret, nil
}

func CreateTokens(s string) []Token {
	rx, _ := regexp.Compile("\r?\n")
	split := rx.Split(s, -1)
	tokens := []Token{}
	for i, line := range split {
		tokens = append(tokens, createTokens(line, i)...)
	}
	return tokens
}

func createTokens(s string, line int) []Token {
	tokens := []Token{}
	queue := runeQueue(s)
	it := 0
	for !queue.Empty() {
		r := queue.Peek()
		switch {
		case isSpace(r):
			it = skipSpace(&queue, it)
		case isNumeric(r):
			token := createToken(&queue, line, it, isNumeric)
			it = token.End() + 1
			tokens = append(tokens, token)
		case isIdentifier(r):
			token := createToken(&queue, line, it, isIdentifier)
			it = token.End() + 1
			tokens = append(tokens, token)
		default:
			token := createSingleToken(&queue, line, it)
			it++
			tokens = append(tokens, token)
		}

	}
	return tokens
}

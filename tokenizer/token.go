package tokenizer

import "fmt"

type token struct {
	begin   int
	content string
}

func (t *token) Begin() int {
	return t.begin
}

func (t *token) End() int {
	return t.begin + len(t.content) - 1
}

func (t *token) Length() int {
	return len(t.content)
}

func (t *token) Content() string {
	return t.content
}

func (t *token) String() string {
	return fmt.Sprintf("Token [begin: %d, end: %d, content: %s]", t.Begin(), t.End(), t.Content())
}

type Token interface {
	Begin() int
	Length() int
	End() int
	Content() string
	String() string
}

func NewToken(begin int, content string) Token {
	return &token{begin: begin, content: content}
}

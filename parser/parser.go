package parser

import (
	"fmt"
	"strconv"

	"github.com/markoczy/ifclib/xp"
	"github.com/markoczy/ifclib/xp/names"
	"github.com/markoczy/ifclib/xp/types"
)

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

func assert(b bool, s string) {
	if !b {
		panic(fmt.Errorf(s))
	}
}

func ParseType(tokens []string) xp.Type {
	var ret xp.Type
	queue := tokenQueue(tokens)
	// var name string
	token := queue.Pop()
	assert(token == "TYPE", "Expected 'TYPE' found "+token)
	name := queue.Pop()
	token = queue.Pop()
	assert(token == "=", "Expected '=' found "+token)

	parent := queue.Pop()
	switch parent {
	case names.Binary:
		ret = parseDerivedNoParams(name, types.Binary, &queue)
	case names.Boolean:
		ret = parseDerivedNoParams(name, types.Boolean, &queue)
	case names.Integer:
		ret = parseDerivedNoParams(name, types.Integer, &queue)
	case names.Logical:
		ret = parseDerivedNoParams(name, types.Logical, &queue)
	case names.Number:
		ret = parseDerivedNoParams(name, types.Number, &queue)
	case names.Real:
		ret = parseDerivedNoParams(name, types.Real, &queue)
	case names.String:
		ret = parseStringDerived(name, &queue)
	case names.Array:
		ret = parseArrayLike(name, &queue, types.NewArray)
	case names.List:
		ret = parseArrayLike(name, &queue, types.NewList)
	case names.Set:
		ret = parseArrayLike(name, &queue, types.NewSet)
	default:
		// TODO
		panic(fmt.Errorf("Unexpected parent type name " + parent))
	}
	token = queue.Pop()
	assert(token == ";", "Expected ';' found "+token)
	// TODO where parser
	return ret
}

func parseDerivedNoParams(name string, parent xp.Type, tokens *tokenQueue) xp.Type {
	return types.NewDerived(name, parent)
}

func parseStringDerived(name string, tokens *tokenQueue) xp.Type {
	var err error
	length := -1
	fixed := false
	if tokens.Peek() == "(" {
		tokens.Pop()
		length, err = strconv.Atoi(tokens.Pop())
		if err != nil {
			panic(fmt.Errorf("Could not parse length to int %w", err))
		}
		tokens.Pop()
		if tokens.Peek() == "FIXED" {
			tokens.Pop()
			fixed = true
		}
	}
	if length == -1 && fixed == false {
		return types.NewDerived(name, types.String)
	}
	return types.NewDerived(name, types.NewString(0, length, fixed))
}

func parseArrayLike(name string, tokens *tokenQueue, generator func(int, int, xp.Type) xp.Type) xp.Type {
	var (
		min, max int
		err      error
	)

	token := tokens.Pop()
	assert(token == "[", "Expected '[' found "+token)

	token = tokens.Pop()
	min, err = strconv.Atoi(token)
	if err != nil {
		panic(fmt.Errorf("Failed to parse min value from token %s %w", token, err))
	}
	token = tokens.Pop()
	assert(token == ":", "Expected ':' found "+token)

	token = tokens.Pop()
	if token == "?" {
		max = -1
	} else {
		max, err = strconv.Atoi(token)
		if err != nil {
			panic(fmt.Errorf("Failed to parse max value from token %s %w", token, err))
		}
	}
	token = tokens.Pop()
	assert(token == "]", "Expected ']' found "+token)
	token = tokens.Pop()
	assert(token == "OF", "Expected 'OF' found "+token)

	// TODO pre-parsed types
	var parent xp.Type
	token = tokens.Pop()
	switch token {
	case names.Binary:
		parent = types.Binary
	case names.Boolean:
		parent = types.Boolean
	case names.Integer:
		parent = types.Integer
	case names.Logical:
		parent = types.Logical
	case names.Number:
		parent = types.Number
	case names.Real:
		parent = types.Real
	case names.String:
		parent = types.String
	default:
		panic(fmt.Errorf("Unexpected parent type name: %s", token))
	}

	return types.NewDerived(name, generator(min, max, parent))
}

func noop(i ...interface{}) {}

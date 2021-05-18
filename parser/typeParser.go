package parser

import (
	"fmt"
	"strconv"

	"github.com/markoczy/ifclib/xp"
	"github.com/markoczy/ifclib/xp/elems"
	"github.com/markoczy/ifclib/xp/names"
	"github.com/markoczy/ifclib/xp/types"
)

func ParseType(tokens []Token, mp elems.Map) xp.Type {
	queue := newTokenQueue(tokens)
	popAndAssertEquals("TYPE", queue)
	name := queue.Pop().Content()
	popAndAssertEquals("=", queue)
	parent := parseType(queue, mp)
	popAndAssertEquals(";", queue)
	// TODO where parser
	return types.NewDerived(name, parent)
}

func parseType(queue *tokenQueue, mp elems.Map) xp.Element {
	var el xp.Element
	parentName := queue.Pop().Content()
	switch parentName {
	case names.Binary:
		el = types.Binary
	case names.Boolean:
		el = types.Boolean
	case names.Integer:
		el = types.Integer
	case names.Logical:
		el = types.Logical
	case names.Number:
		el = types.Number
	case names.Real:
		el = types.Real
	case names.String:
		el = parseString(queue)
	case names.Array:
		el = parseArrayLike(queue, types.NewArray, mp)
	case names.List:
		el = parseArrayLike(queue, types.NewList, mp)
	case names.Set:
		el = parseArrayLike(queue, types.NewSet, mp)
	case names.Enumeration:
		el = parseEnumeration(queue)
	default:
		el = mp.Lookup(parentName)
	}
	return el
}

// func parseDerivedNoParams(name string, parent xp.Type, tokens *tokenQueue) xp.Type {
// 	return types.NewDerived(name, parent)
// }

func parseString(tokens *tokenQueue) xp.Type {
	var err error
	length := -1
	fixed := false
	if tokens.Peek().Content() == "(" {
		tokens.Pop()
		length, err = strconv.Atoi(tokens.Pop().Content())
		if err != nil {
			panic(fmt.Errorf("Could not parse length to int %w", err))
		}
		tokens.Pop()
		if tokens.Peek().Content() == "FIXED" {
			tokens.Pop()
			fixed = true
		}
	}
	if length == -1 && fixed == false {
		return types.String
	}
	return types.NewString(0, length, fixed)
}

func parseArrayLike(queue *tokenQueue, generator func(int, int, xp.Element) xp.Type, mp elems.Map) xp.Type {
	var (
		min, max int
		err      error
		token    Token
	)

	popAndAssertEquals("[", queue)
	token = queue.Pop()
	min, err = strconv.Atoi(token.Content())
	if err != nil {
		panic(fmt.Errorf("Failed to parse min value from token %s %w", token, err))
	}
	popAndAssertEquals(":", queue)

	token = queue.Pop()
	if token.Content() == "?" {
		max = -1
	} else {
		max, err = strconv.Atoi(token.Content())
		if err != nil {
			panic(fmt.Errorf("Failed to parse max value from token %s %w", token, err))
		}
	}
	popAndAssertEquals("]", queue)
	popAndAssertEquals("OF", queue)

	var parent xp.Element
	token = queue.Pop()
	switch token.Content() {
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
		parent = mp.Lookup(token.Content())
	}
	return generator(min, max, parent)
}

func parseEnumeration(queue *tokenQueue) xp.Type {
	popAndAssertEquals("OF", queue)
	popAndAssertEquals("(", queue)
	names := []string{}
	for queue.Peek().Content() != ")" {
		names = append(names, queue.Pop().Content())
		if queue.Peek().Content() == "," {
			queue.Pop()
		}
	}
	queue.Pop()
	return types.NewEnumeration(names)
}

// func parseSelect(name string, tokens *tokenQueue, mp types.TypeMap) xp.Type {
// 	oneOf := []xp.Type{}
// 	token := tokens.Pop()
// 	assert(token == "(", "Expected '(' found "+token)
// 	for tokens.Peek() != ")" {
// 		// names = append(names, tokens.Pop())
// 		name := tokens.Pop()
// 		tp := mp.Lookup(name)
// 		if tp == nil {
// 			panic(fmt.Errorf("Unresolvable type in SELECT: " + name))
// 		}
// 		oneOf = append(oneOf, tp)

// 		if tokens.Peek() == "," {
// 			tokens.Pop()
// 		}
// 	}
// 	tokens.Pop()
// 	return types.NewSelect(oneOf)
// }

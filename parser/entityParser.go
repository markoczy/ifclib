package parser

import (
	"github.com/markoczy/ifclib/xp"
	"github.com/markoczy/ifclib/xp/elems"
	"github.com/markoczy/ifclib/xp/types"
)

func ParseEntity(tokens []string, mp elems.Map) xp.Entity {
	//
	var (
		abstract    bool
		supertypeOf []xp.Element
		subtypeOf   xp.Element
		inverseAttr []xp.InverseAttr
		properties  []xp.Property
	)
	queue := newTokenQueue(tokens)

	popAndAssertEquals("ENTITY", queue)
	for queue.Peek() != "END_ENTITY" {
		switch queue.Peek() {
		case "SUBTYPE":
			parseSubtypeOf(queue, mp)
		case "ABSTRACT", "SUPERTYPE":
			parseSupertypeOf(queue, mp)
		case "INVERSE":
			parseInverse(queue, mp)
		case "WHERE":
			parseWhere(queue)
		default:
			parseProperties(queue, mp)
		}
	}
	return types.NewDefaultEntity(abstract, supertypeOf, subtypeOf, inverseAttr, properties)
}

func parseSubtypeOf(queue *tokenQueue, mp elems.Map) xp.Element {
	popAndAssertEquals("SUBTYPE", queue)
	popAndAssertEquals("OF", queue)
	popAndAssertEquals("(", queue)
	token := queue.Pop()
	subtypeOf := mp.Lookup(token)
	token = queue.Pop()
	popAndAssertEquals(")", queue)
	if queue.Peek() == ";" {
		queue.Pop()
	}
	return subtypeOf
}

func parseSupertypeOf(queue *tokenQueue, mp elems.Map) (bool, []xp.Element) {
	abstract := false
	supertypeOf := []xp.Element{}
	token := queue.Peek()
	if token == "ABSTRACT" {
		abstract = true
		queue.Pop()
	}
	popAndAssertEquals("SUPERTYPE", queue)
	popAndAssertEquals("OF", queue)
	popAndAssertEquals("(", queue)
	//? Only found supertype with ONEOF, maybe in other step files this is optional
	popAndAssertEquals("ONEOF", queue)
	popAndAssertEquals("(", queue)
	for queue.Peek() != ")" {
		name := queue.Pop()
		supertypeOf = append(supertypeOf, mp.Lookup(name))
		if queue.Peek() == "," {
			queue.Pop()
		}
	}
	popAndAssertEquals(")", queue)
	popAndAssertEquals(")", queue)
	return abstract, supertypeOf
}

func parseProperties(queue *tokenQueue, mp elems.Map) []xp.Property {
	ret := []xp.Property{}
	//! no better exit condition yet..
	for queue.Peek() != "WHERE" && queue.Peek() != "END_ENTITY" {
		optional := false
		if queue.Peek() == "OPTIONAL" {
			optional = true
			queue.Pop()
		}
		name := queue.Pop()
		el := parseType(queue, mp)
		ret = append(ret, types.NewDefaultProperty(name, el, optional))
	}
	return ret
}

func parseInverse(queue *tokenQueue, mp elems.Map) []xp.InverseAttr {
	ret := []xp.InverseAttr{}
	popAndAssertEquals("INVERSE", queue)
	// Unfortunately no better end condition
	for queue.Peek() != "WHERE" && queue.Peek() != "END_ENTITY" {
		name := queue.Pop()
		popAndAssertEquals(":", queue)
		tp := parseType(queue, mp)
		popAndAssertEquals("FOR", queue)
		propName := queue.Pop()
		ret = append(ret, types.NewDefaultInverseAttr(name, tp, propName))
		popAndAssertEquals(";", queue)
	}
	return ret
}

func parseWhere(queue *tokenQueue) {
	// TODO We just skip the WHERE Statements for now...
	popAndAssertEquals("WHERE", queue)
	for queue.Peek() != ";" {
		queue.Pop()
	}
}

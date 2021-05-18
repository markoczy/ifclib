package parser

import (
	"github.com/markoczy/ifclib/xp"
	"github.com/markoczy/ifclib/xp/elems"
	"github.com/markoczy/ifclib/xp/types"
)

func ParseEntity(tokens []Token, mp elems.Map) xp.Entity {
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
	name := queue.Pop().Content()
	for queue.Peek().Content() != "END_ENTITY" {
		switch queue.Peek().Content() {
		case "SUBTYPE":
			subtypeOf = parseSubtypeOf(queue, mp)
		case "ABSTRACT", "SUPERTYPE":
			abstract, supertypeOf = parseSupertypeOf(queue, mp)
		case "INVERSE":
			inverseAttr = parseInverse(queue, mp)
		case "WHERE":
			parseWhere(queue)
		default:
			properties = parseProperties(queue, mp)
		}
	}
	ret := types.NewDefaultEntity(name, abstract, supertypeOf, subtypeOf, inverseAttr, properties)
	// fmt.Println("*** Parsed Entity:", ret)
	return ret
}

func parseSubtypeOf(queue *tokenQueue, mp elems.Map) xp.Element {
	popAndAssertEquals("SUBTYPE", queue)
	popAndAssertEquals("OF", queue)
	popAndAssertEquals("(", queue)
	token := queue.Pop()
	subtypeOf := mp.Lookup(token.Content())
	// token = queue.Pop()
	popAndAssertEquals(")", queue)
	if queue.Peek().Content() == ";" {
		queue.Pop()
	}
	return subtypeOf
}

func parseSupertypeOf(queue *tokenQueue, mp elems.Map) (bool, []xp.Element) {
	abstract := false
	supertypeOf := []xp.Element{}
	token := queue.Peek()
	if token.Content() == "ABSTRACT" {
		abstract = true
		queue.Pop()
	}
	popAndAssertEquals("SUPERTYPE", queue)
	popAndAssertEquals("OF", queue)
	popAndAssertEquals("(", queue)
	//? Only found supertype with ONEOF, maybe in other step files this is optional
	popAndAssertEquals("ONEOF", queue)
	popAndAssertEquals("(", queue)
	for queue.Peek().Content() != ")" {
		name := queue.Pop().Content()
		supertypeOf = append(supertypeOf, mp.Lookup(name))
		if queue.Peek().Content() == "," {
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
	for queue.Peek().Content() != "WHERE" && queue.Peek().Content() != "END_ENTITY" {
		optional := false
		if queue.Peek().Content() == "OPTIONAL" {
			optional = true
			queue.Pop()
		}
		name := queue.Pop().Content()
		el := parseType(queue, mp)
		ret = append(ret, types.NewDefaultProperty(name, el, optional))
	}
	return ret
}

func parseInverse(queue *tokenQueue, mp elems.Map) []xp.InverseAttr {
	ret := []xp.InverseAttr{}
	popAndAssertEquals("INVERSE", queue)
	// Unfortunately no better end condition
	for queue.Peek().Content() != "WHERE" && queue.Peek().Content() != "END_ENTITY" {
		name := queue.Pop().Content()
		popAndAssertEquals(":", queue)
		tp := parseType(queue, mp)
		popAndAssertEquals("FOR", queue)
		propName := queue.Pop().Content()
		ret = append(ret, types.NewDefaultInverseAttr(name, tp, propName))
		popAndAssertEquals(";", queue)
	}
	return ret
}

func parseWhere(queue *tokenQueue) {
	// TODO We just skip the WHERE Statements for now...
	popAndAssertEquals("WHERE", queue)
	for queue.Peek().Content() != ";" {
		queue.Pop()
	}
}

// TODO Parse DERIVE, UNIQUE

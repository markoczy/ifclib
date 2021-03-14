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
	queue := tokenQueue(tokens)

	// var name string
	token := queue.Pop()
	assert(token == "ENTITY", "Expected 'ENTITY' found "+token)

	token = queue.Peek()
	switch token {
	case "SUBTYPE":

	case "ABSTRACT", "SUPERTYPE":

	case "INVERSE":

	// default:
	case "WHERE":
		parseWhere(&queue)
	}
	// name := queue.Pop()
	// token = queue.Pop()
	// assert(token == "=", "Expected '=' found "+token)
	noop(supertypeOf)

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

func parseInverse(queue *tokenQueue, mp elems.Map) []xp.InverseAttr {
	ret := []xp.InverseAttr{}
	popAndAssertEquals("INVERSE", queue)
	// Unfortunately no better end condition
	for queue.Peek() != "WHERE" && queue.Peek() != "END_ENTITY" {
		// name := queue.Pop()

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

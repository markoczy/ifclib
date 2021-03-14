package types

import (
	"fmt"

	"github.com/markoczy/ifclib/xp"
)

type defaultProperty struct {
	name     string
	element  xp.Element
	optional bool
}

func (p *defaultProperty) Name() string {
	return p.name
}

func (p *defaultProperty) Element() xp.Element {
	return p.element
}

func (p *defaultProperty) Optional() bool {
	return p.optional
}

func (p *defaultProperty) String() string {
	return fmt.Sprintf("defaultProperty: { name: %s, element: %v, optional: %v }", p.name, p.element, p.optional)
}

func NewDefaultProperty(name string, element xp.Element, optional bool) xp.Property {
	return &defaultProperty{
		name:     name,
		element:  element,
		optional: optional,
	}
}

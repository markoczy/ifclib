package types

import (
	"fmt"

	"github.com/markoczy/ifclib/xp"
)

type defaultInverseAttr struct {
	name        string
	element     xp.Element
	forProperty string
}

func (ia *defaultInverseAttr) Name() string {
	return ia.name
}

func (ia *defaultInverseAttr) Element() xp.Element {
	return ia.element
}

func (ia *defaultInverseAttr) ForProperty() string {
	return ia.forProperty
}

func (ia *defaultInverseAttr) String() string {
	return fmt.Sprintf("defaultInverseAttr: { name: %s, element: %v, forProperty: %s }", ia.name, ia.element, ia.forProperty)
}

func NewDefaultInverseAttr(name string, element xp.Element, forProperty string) xp.InverseAttr {
	return &defaultInverseAttr{
		name:        name,
		element:     element,
		forProperty: forProperty,
	}
}

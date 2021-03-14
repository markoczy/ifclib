package types

import (
	"fmt"

	"github.com/markoczy/ifclib/xp"
)

func NewDefaultEntity(
	abstract bool,
	supertypeOf []xp.Element,
	subtypeOf xp.Element,
	inverse []xp.InverseAttr,
	properties []xp.Property) xp.Entity {

	return &defaultEntity{
		abstract:    abstract,
		supertypeOf: supertypeOf,
		subtypeOf:   subtypeOf,
		inverse:     inverse,
		properties:  properties,
	}
}

type defaultEntity struct {
	abstract    bool
	supertypeOf []xp.Element
	subtypeOf   xp.Element
	inverse     []xp.InverseAttr
	properties  []xp.Property
}

func (e *defaultEntity) Abstract() bool {
	return e.abstract
}

func (e *defaultEntity) SupertypeOf() []xp.Element {
	return e.supertypeOf
}

func (e *defaultEntity) SubtypeOf() xp.Element {
	return e.subtypeOf
}

func (e *defaultEntity) Inverse() []xp.InverseAttr {
	return e.inverse
}

func (e *defaultEntity) Properties() []xp.Property {
	return e.properties
}

func (e *defaultEntity) Type() xp.Type {
	return nil
}

func (e *defaultEntity) Entity() xp.Entity {
	return e
}

func (e *defaultEntity) String() string {
	return fmt.Sprintf("defaultEntity: { abstract: %v, supertypeOf: %v, subtypeOf: %v, inverse: %v, properties: %v }", e.abstract, e.supertypeOf, e.subtypeOf, e.inverse, e.properties)
}

package types

import (
	"fmt"
	"strings"

	"github.com/markoczy/ifclib/xp"
)

func NewDefaultEntity(
	name string,
	abstract bool,
	supertypeOf []xp.Element,
	subtypeOf xp.Element,
	inverse []xp.InverseAttr,
	properties []xp.Property) xp.Entity {

	return &defaultEntity{
		name:        name,
		abstract:    abstract,
		supertypeOf: supertypeOf,
		subtypeOf:   subtypeOf,
		inverse:     inverse,
		properties:  properties,
	}
}

type defaultEntity struct {
	name        string
	abstract    bool
	supertypeOf []xp.Element
	subtypeOf   xp.Element
	inverse     []xp.InverseAttr
	properties  []xp.Property
}

func (e *defaultEntity) Name() string {
	return e.name
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
	supertypeOf := []string{}
	if e.supertypeOf != nil {
		for _, v := range e.supertypeOf {
			supertypeOf = append(supertypeOf, v.Name())
		}
	}
	subtypeOf := "<nil>"
	if e.subtypeOf != nil {
		subtypeOf = e.subtypeOf.Name()
	}
	return fmt.Sprintf("defaultEntity: { abstract: %v, supertypeOf: %v, subtypeOf: %v, inverse: %v, properties: %v }", e.abstract, "["+strings.Join(supertypeOf, ", ")+"]", subtypeOf, e.inverse, e.properties)
}

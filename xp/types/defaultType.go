package types

import (
	"fmt"

	"github.com/markoczy/ifclib/xp"
)

type defaultType struct {
	name       string
	parent     xp.Type
	primitive  bool
	properties []xp.Property
	values     []string
	elements   []xp.Type
	min        int
	max        int
	fixed      bool
}

func newDefaultType(name string, opts ...func(*defaultType)) xp.Type {
	ret := &defaultType{
		name:      name,
		primitive: true,
		min:       0,
		max:       -1,
	}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

func (t *defaultType) Parent() xp.Type {
	return t.parent
}

func (t *defaultType) Name() string {
	return t.name
}

func (t *defaultType) Primitive() bool {
	return t.primitive
}

func (t *defaultType) Properties() []xp.Property {
	return t.properties
}

func (t *defaultType) Values() []string {
	return t.values
}

func (t *defaultType) Elements() []xp.Type {
	return t.elements
}

func (t *defaultType) Min() int {
	return t.min
}

func (t *defaultType) Max() int {
	return t.max
}

func (t *defaultType) Fixed() bool {
	return t.fixed
}

func (t *defaultType) String() string {
	return fmt.Sprintf("defaultType: { parent: %v, name: %v, primitive: %v, properties: %v, values: %v, elements: %v, min: %v, max: %v, fixed: %v }", t.parent, t.name, t.primitive, t.properties, t.values, t.elements, t.min, t.max, t.fixed)
}

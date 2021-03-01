package builtin

import "github.com/markoczy/ifclib/xp"

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
		fixed:     true,
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

package types

import (
	"fmt"
	"strings"

	"github.com/markoczy/ifclib/xp"
)

type TypeMap interface {
	Lookup(name string) xp.Type
	Assign(name string, tp xp.Type)
	String() string
}

type typeMap map[string]*mappedType

func NewTypeMap(names []string) TypeMap {
	ret := typeMap(map[string]*mappedType{})
	for _, name := range names {
		ret[name] = &mappedType{}
	}
	return &ret
}

func (m *typeMap) Lookup(name string) xp.Type {
	return (*m)[name]
}

func (m *typeMap) Assign(name string, tp xp.Type) {
	(*m)[name].tp = tp
}

func (m *typeMap) String() string {
	sb := strings.Builder{}
	for k, v := range *(*map[string]*mappedType)(m) {
		sb.WriteString(fmt.Sprintf("%s -> %v\n\n", k, v))
	}
	return sb.String()
}

type mappedType struct {
	tp xp.Type
}

func (t *mappedType) Parent() xp.Type {
	return t.tp.Parent()
}

func (t *mappedType) Name() string {
	return t.tp.Name()
}

func (t *mappedType) Primitive() bool {
	return t.tp.Primitive()
}

func (t *mappedType) Properties() []xp.Property {
	return t.tp.Properties()
}

func (t *mappedType) Values() []string {
	return t.tp.Values()
}

func (t *mappedType) Elements() []xp.Type {
	return t.tp.Elements()
}

func (t *mappedType) Min() int {
	return t.tp.Min()
}

func (t *mappedType) Max() int {
	return t.tp.Max()
}

func (t *mappedType) Fixed() bool {
	return t.tp.Fixed()
}

func (t *mappedType) String() string {
	return fmt.Sprintf("mappedType: { tp: %v }", t.tp)
}

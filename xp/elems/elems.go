package elems

import (
	"fmt"
	"strings"

	"github.com/markoczy/ifclib/xp"
)

type Map interface {
	Lookup(name string) xp.Element
	Assign(name string, tp xp.Element)
	String() string
}

func NewMap(names []string) Map {
	ret := elemMap(map[string]*mappedElem{})
	for _, name := range names {
		ret[name] = &mappedElem{}
	}
	return &ret
}

type elemMap map[string]*mappedElem

func (m *elemMap) Lookup(name string) xp.Element {
	return (*m)[name]
}

func (m *elemMap) Assign(name string, elem xp.Element) {
	(*m)[name].elem = elem
}

func (m *elemMap) String() string {
	sb := strings.Builder{}
	for k, v := range *(*map[string]*mappedElem)(m) {
		sb.WriteString(fmt.Sprintf("%s -> %v\n\n", k, v))
	}
	return sb.String()
}

type mappedElem struct {
	elem xp.Element
}

func (el *mappedElem) Name() string {
	return el.elem.Name()
}

func (el *mappedElem) Type() xp.Type {
	return el.elem.Type()
}

func (el *mappedElem) Entity() xp.Entity {
	return el.elem.Entity()
}

func (t *mappedElem) String() string {
	return fmt.Sprintf("mappedElem: { elem: %v }", t.elem)
}

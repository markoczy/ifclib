package types

import (
	"github.com/markoczy/ifclib/xp"
	"github.com/markoczy/ifclib/xp/names"
)

var (
	Binary  = newDefaultType(names.Binary)
	Boolean = newDefaultType(names.Boolean)
	Integer = newDefaultType(names.Integer)
	Logical = newDefaultType(names.Logical)
	Number  = newDefaultType(names.Number)
	Real    = newDefaultType(names.Real)
	// String with no length restrictions
	String = newDefaultType(names.String)
)

func NewString(min, max int, fixed bool) xp.Type {
	return newDefaultType(names.String, func(dt *defaultType) {
		dt.min = min
		dt.max = max
		dt.fixed = fixed
	})
}

func NewEnumeration(name string, values []string) xp.Type {
	// Enumerations are always a derived type
	return NewDerived(name, newDefaultType(names.Enumeration, func(dt *defaultType) {
		dt.values = values
	}))
}

func NewArray(min, max int, of xp.Type) xp.Type {
	return newDefaultType(names.Array, func(dt *defaultType) {
		dt.min = min
		dt.max = max
		dt.elements = []xp.Type{of}
	})
}

func NewList(min, max int, of xp.Type) xp.Type {
	// ? difference between array, list and set?
	return newDefaultType(names.List, func(dt *defaultType) {
		dt.min = min
		dt.max = max
		dt.elements = []xp.Type{of}
	})
}

func NewSet(min, max int, of xp.Type) xp.Type {
	// ? difference between array, list and set?
	return newDefaultType(names.Set, func(dt *defaultType) {
		dt.min = min
		dt.max = max
		dt.elements = []xp.Type{of}
	})
}

func NewSelect(oneOf []xp.Type) xp.Type {
	return newDefaultType(names.Select, func(dt *defaultType) {
		dt.elements = oneOf
	})
}

func NewDerived(name string, parent xp.Type) xp.Type {
	return newDefaultType(name, func(dt *defaultType) {
		dt.parent = parent
		dt.primitive = false
	})
}

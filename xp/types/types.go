package types

import "github.com/markoczy/ifclib/xp"

var (
	Binary  = newDefaultType("BINARY")
	Boolean = newDefaultType("BOOLEAN")
	Integer = newDefaultType("INTEGER")
	Logical = newDefaultType("LOGICAL")
	Number  = newDefaultType("NUMBER")
	Real    = newDefaultType("REAL")
	// String with no length restrictions
	String = newDefaultType("STRING")
)

func NewString(min, max int, fixed bool) xp.Type {
	return newDefaultType("STRING", func(dt *defaultType) {
		dt.min = min
		dt.max = max
		dt.fixed = fixed
	})
}

func NewEnumeration(name string, values []string) xp.Type {
	// Enumerations are always a derived type
	return NewDerived(name, newDefaultType("ENUMERATION", func(dt *defaultType) {
		dt.values = values
	}))
}

func NewArray(min, max int, of xp.Type) xp.Type {
	return newDefaultType("ARRAY", func(dt *defaultType) {
		dt.min = min
		dt.max = max
		dt.elements = []xp.Type{of}
	})
}

func NewList(min, max int, of xp.Type) xp.Type {
	// ? difference between array, list and set?
	return newDefaultType("LIST", func(dt *defaultType) {
		dt.min = min
		dt.max = max
		dt.elements = []xp.Type{of}
	})
}

func NewSet(min, max int, of xp.Type) xp.Type {
	// ? difference between array, list and set?
	return newDefaultType("SET", func(dt *defaultType) {
		dt.min = min
		dt.max = max
		dt.elements = []xp.Type{of}
	})
}

func NewSelect(oneOf []xp.Type) xp.Type {
	return newDefaultType("SELECT", func(dt *defaultType) {
		dt.elements = oneOf
	})
}

func NewDerived(name string, parent xp.Type) xp.Type {
	return newDefaultType(name, func(dt *defaultType) {
		dt.parent = parent
		dt.primitive = false
	})
}

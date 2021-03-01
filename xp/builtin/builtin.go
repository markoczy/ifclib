package builtin

import "github.com/markoczy/ifclib/xp"

var (
	Binary  = newDefaultType("BINARY")
	Boolean = newDefaultType("BOOLEAN")
	Integer = newDefaultType("INTEGER")
	Logical = newDefaultType("LOGICAL")
	Number  = newDefaultType("NUMBER")
	Real    = newDefaultType("REAL")
)

func NewString(min, max int, fixed bool) xp.Type {
	return newDefaultType("STRING", func(dt *defaultType) {
		dt.min = min
		dt.max = max
		dt.fixed = fixed
	})
}

func NewEnumeration(values []string) xp.Type {
	return newDefaultType("ENUMERATION", func(dt *defaultType) {
		dt.values = values
	})
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

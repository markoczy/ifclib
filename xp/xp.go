package xp

type Property interface {
	Name() string
	Type() Type
}

type Type interface {
	// Parent Type "is a" relationship
	Parent() Type
	// Identifier of the type
	Name() string
	// Wether the type is an Express primitive
	Primitive() bool
	// Named Properties
	Properties() []Property
	// Enum Constants
	Values() []string
	// Arrays, Lists: The Type of children elements (only 1 allowed)
	//
	// Select: The possible types ("one of")
	Elements() []Type
	// For strings, arrays etc.
	Min() int
	// For strings, arrays etc.
	Max() int
	// For Strings, if true value Max counts as the limit
	Fixed() bool
	// Stringer interface
	String() string
}

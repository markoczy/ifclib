package xp

// Element the common interface for any Type or Entity
type Element interface {
	Name() string
	// Type returns the Type representation or nil if Element is not a Type
	Type() Type
	// Entity returns the Entity representation or nil if Element is not an Entity
	Entity() Entity
}

type Property interface {
	Name() string
	Element() Element
	Optional() bool
	// Stringer interface
	String() string
}

type Type interface {
	// Parent Type "is a" relationship
	Parent() Element
	// Identifier of the type
	Name() string
	// Wether the type is an Express primitive
	Primitive() bool
	// Enum Constants
	Values() []string
	// Arrays, Lists: The Type of children elements (only 1 allowed)
	//
	// Select: The possible types ("one of")
	Elements() []Element
	// For strings, arrays etc.
	Min() int
	// For strings, arrays etc.
	Max() int
	// For Strings, if true value Max counts as the limit
	Fixed() bool
	// Element interface
	Type() Type
	// Element interface
	Entity() Entity
	// Stringer interface
	String() string
	//
	// TODO WHERE
	//
}

type InverseAttr interface {
	Name() string
	Element() Element
	ForProperty() string
	// Stringer interface
	String() string
}

type Entity interface {
	Name() string
	Abstract() bool
	SupertypeOf() []Element
	SubtypeOf() Element
	Inverse() []InverseAttr
	Properties() []Property
	// Element interface
	Type() Type
	// Element interface
	Entity() Entity
	// Stringer interface
	String() string
	//
	// TODO DERIVE WHERE
	//
}

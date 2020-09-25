package object

import "fmt"

const (
	integerObj = "INTEGER"
	booleanObj = "BOOLEAN"
	nullObj    = "NULL"
)

//ObjectType is an enum that represents the object type
type ObjectType string

//Object is the base object in monkey
type Object interface {
	// Type gets the ObjectType
	Type() ObjectType

	//Inspect gets the string representation
	Inspect() string
}

//Integer is the basic number object
type Integer struct {
	Value int64
}

// Type gets the ObjectType
func (i *Integer) Type() ObjectType { return integerObj }

//Inspect gets the string representation
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

//Boolean is an object representing true or false
type Boolean struct {
	Value bool
}

// Type gets the ObjectType
func (b *Boolean) Type() ObjectType { return booleanObj }

//Inspect gets the string representation
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

//Null is an object representing absence of a value
type Null struct{}

// Type gets the ObjectType
func (n *Null) Type() ObjectType { return nullObj }

//Inspect gets the string representation
func (n *Null) Inspect() string { return "null" }

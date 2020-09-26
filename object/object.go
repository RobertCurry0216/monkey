package object

import "fmt"

const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
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
func (i *Integer) Type() ObjectType { return IntegerObj }

//Inspect gets the string representation
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

//Boolean is an object representing true or false
type Boolean struct {
	Value bool
}

// Type gets the ObjectType
func (b *Boolean) Type() ObjectType { return BooleanObj }

//Inspect gets the string representation
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

//Null is an object representing absence of a value
type Null struct{}

// Type gets the ObjectType
func (n *Null) Type() ObjectType { return NullObj }

//Inspect gets the string representation
func (n *Null) Inspect() string { return "null" }

//ReturnValue is pased around the evaluator to determine the value to be returned
type ReturnValue struct {
	Value Object
}

// Type gets the ObjectType
func (rv *ReturnValue) Type() ObjectType { return ReturnValueObj }

//Inspect gets the string representation
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

//Error is an user error
type Error struct {
	Message string
}

// Type gets the ObjectType
func (e *Error) Type() ObjectType { return ErrorObj }

//Inspect gets the string representation
func (e *Error) Inspect() string { return "Error: " + e.Message }

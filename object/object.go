package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"monkey/ast"
	"strings"
)

const (
	IntegerObj     = "INTEGER"
	StringObj      = "STRING"
	BooleanObj     = "BOOLEAN"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
	FunctionObj    = "FUNCTION"
	BuiltinObj     = "BUILTIN"
	ArrayObj       = "ARRAY"
	HashObj        = "HASH"
)

//ObjectType is an enum that represents the object type
type ObjectType string

//BuiltinFunction is for the functions that come with the monkey language
type BuiltinFunction func(args ...Object) Object

//Object is the base object in monkey
type Object interface {
	// Type gets the ObjectType
	Type() ObjectType

	//Inspect gets the string representation
	Inspect() string
}

//Hashable represents types that can be used as a HashKey
type Hashable interface {
	HashKey() HashKey
}

//HashKey is for hashing values
type HashKey struct {
	Type  ObjectType
	Value uint64
}

//Integer is the basic number object
type Integer struct {
	Value int64
}

// Type gets the ObjectType
func (i *Integer) Type() ObjectType { return IntegerObj }

//Inspect gets the string representation
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

//HashKey gets a unique value for this object
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

//String is the string primative
type String struct {
	Value string
}

// Type gets the ObjectType
func (s *String) Type() ObjectType { return StringObj }

//Inspect gets the string representation
func (s *String) Inspect() string { return s.Value }

//HashKey gets a unique value for this object
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

//Boolean is an object representing true or false
type Boolean struct {
	Value bool
}

// Type gets the ObjectType
func (b *Boolean) Type() ObjectType { return BooleanObj }

//Inspect gets the string representation
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

//HashKey gets a unique value for this object
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

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

//Function is a function
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Enviroment
}

// Type gets the ObjectType
func (f *Function) Type() ObjectType { return FunctionObj }

//Inspect gets the string representation
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") \n{")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// Builtin is a function that comes with monkey
type Builtin struct {
	Fn BuiltinFunction
}

// Type gets the ObjectType
func (bi *Builtin) Type() ObjectType { return BuiltinObj }

//Inspect gets the string representation
func (bi *Builtin) Inspect() string { return "builtin function" }

// Array is a function that comes with monkey
type Array struct {
	Elements []Object
}

// Type gets the ObjectType
func (a *Array) Type() ObjectType { return ArrayObj }

//Inspect gets the string representation
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashPair is a key value pair used in a hashmap
type HashPair struct {
	Key   Object
	Value Object
}

//Hash is a hashmap
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type gets the ObjectType
func (h *Hash) Type() ObjectType { return HashObj }

//Inspect gets the string representation
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

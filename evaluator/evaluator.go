package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

//Eval evaluates a node and returns an object
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

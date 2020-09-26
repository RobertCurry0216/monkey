package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

var (
	trueObj  = &object.Boolean{Value: true}
	falseObj = &object.Boolean{Value: false}
	nullObj  = &object.Null{}
)

//Eval evaluates a node and returns an object
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	//Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
		//Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBoolObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfIxExpression(node.Operator, left, right)
	}

	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}

func nativeBoolToBoolObject(input bool) *object.Boolean {
	if input {
		return trueObj
	}
	return falseObj
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		return nullObj
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case trueObj:
		return falseObj
	case falseObj:
		return trueObj
	case nullObj:
		return trueObj
	default:
		return falseObj
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.IntegerObj {
		return nullObj
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalInfIxExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		return evalIntegerInfixExpression(operator, left, right)

		// TODO: determine truthyness
	case operator == "==":
		return nativeBoolToBoolObject(left == right)
	case operator == "!=":
		return nativeBoolToBoolObject(left != right)
	default:
		return nullObj
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBoolObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBoolObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBoolObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBoolObject(leftVal != rightVal)
	default:
		return nullObj
	}
}

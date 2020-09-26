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
		return evalProgram(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.BlockStatement:
		return evalBlockStatement(node.Statements)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
		return &object.ReturnValue{Value: val}
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
	case *ast.IfExpression:
		return evalIfExpresssion(node)
	}
	return nil
}

func evalProgram(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}

func evalBlockStatement(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
		if result != nil && result.Type() == object.ReturnValueObj {
			return result
		}
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
		return nativeBoolToBoolObject(isTruthy(left) == isTruthy(right))
	case operator == "!=":
		return nativeBoolToBoolObject(isTruthy(left) != isTruthy(right))
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

func evalIfExpresssion(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)

	if isTruthy(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return nullObj
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case trueObj:
		return true
	case falseObj:
		return false
	case nullObj:
		return false
	default:
		switch {
		case obj.Type() == object.IntegerObj:
			return obj.(*object.Integer).Value != 0
		default:
			return false
		}
	}
}

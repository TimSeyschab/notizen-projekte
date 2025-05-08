package eval

import (
	"fmt"
	"interpreter/ast"
	"interpreter/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		if isError(left) {
			return left
		}
		right := Eval(node.Right)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)

	case *ast.BlockStatement:
		return evalBlockStatement(node)

	case *ast.IfExpression:
		ifCond := Eval(node.Condition)
		if isError(ifCond) {
			return ifCond
		}

		if isTruthy(ifCond) {
			return Eval(node.Consequence)
		} else if node.Alternative != nil {
			return Eval(node.Alternative)
		}

		return NULL

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return getNativeBooleanObject(node.Value)

	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	}

	return NULL
}

func isTruthy(ifCond object.Object) bool {
	switch ifCond {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case isNumber(left) && isNumber(right):
		return evalNumberInfixExpression(operator, left, right)
	case operator == "==":
		return getNativeBooleanObject(left == right)
	case operator == "!=":
		return getNativeBooleanObject(left != right)
	case left.Type() != right.Type():
		return createError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return createError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func isNumber(obj object.Object) bool {
	switch obj.Type() {
	case object.INTEGER_OBJ, object.FLOAT_OBJ:
		return true
	default:
		return false
	}
}

func evalNumberInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal, isLeftFloat := getValueAndType(left)
	rightVal, isRightFloat := getValueAndType(right)
	if leftVal == nil || rightVal == nil {
		return NULL
	}

	switch operator {
	case "<":
		return getNativeBooleanObject(leftVal.(float64) < rightVal.(float64))
	case ">":
		return getNativeBooleanObject(leftVal.(float64) > rightVal.(float64))
	case "==":
		return getNativeBooleanObject(leftVal.(float64) == rightVal.(float64))
	case "!=":
		return getNativeBooleanObject(leftVal.(float64) != rightVal.(float64))
	}

	if isLeftFloat || isRightFloat {
		return evalFloatInfix(operator, leftVal.(float64), rightVal.(float64))
	}
	return evalIntInfix(operator, int64(leftVal.(float64)), int64(rightVal.(float64)))
}

func getValueAndType(obj object.Object) (interface{}, bool) {
	switch v := obj.(type) {
	case *object.Integer:
		return float64(v.Value), false
	case *object.Float:
		return v.Value, true
	default:
		return nil, false
	}
}

func evalFloatInfix(operator string, leftVal, rightVal float64) object.Object {
	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	default:
		return createError("unknown operator: %f %s %f", leftVal, operator, rightVal)
	}
}

func evalIntInfix(operator string, leftVal, rightVal int64) object.Object {
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	default:
		return createError("unknown operator: %d %s %d", leftVal, operator, rightVal)
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		return createError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object {
	switch right := right.(type) {
	case *object.Float:
		right.Value = right.Value * -1
		return right
	case *object.Integer:
		right.Value = right.Value * -1
		return right
	default:
		return createError("unknown operator: -%s", right.Type())
	}

	return NULL
}

func getNativeBooleanObject(b bool) *object.Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement)

		// check for ReturnValue or else last Statement will be result
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func createError(format string, a ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
